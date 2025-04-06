package batch

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

const (
	DefaultMaxQueueSize       = 2048
	DefaultExportTimeout      = 30000
	DefaultMaxExportBatchSize = 512
)

type Processor interface {
	OnStart(parent context.Context, data Data)
	OnEnd(s Data)
	Shutdown(ctx context.Context) error
	ForceFlush(ctx context.Context) error
}

type Data struct {
}

type SpanExporter interface {
	ExportSpans(ctx context.Context, data []Data) error
}

// BatchSpanProcessorOptions 用来配置我们的处理器
type BatchSpanProcessorOptions struct {
	//设置当前批次对应的超时时间
	BatchTimeout time.Duration
	//设置导出的时间
	ExportTimeout      time.Duration
	MaxExportBatchSize int
	MaxQueueSize       int
	BlockOnQueueFull   bool
}

type batchDataProcessor struct {
	e SpanExporter

	//这里我们需要设计一个导出器 这个导出器可以是从数据库中抽取一批数据
	//也可以是通过一个rpc 调用, 或者是向消息队列去发送对应的消息
	o     BatchSpanProcessorOptions
	queue chan Data
	// 表示有多少 span 其实是被丢弃的
	dropped uint32
	// 存储的数据
	batch []Data
	// 用来建造对应的临界区
	batchMutex sync.Mutex
	// 这个字段主要用来完成我们整个batch 的定时控制
	timer *time.Timer
	// 下面这些字段主要是用来处理这个 processor 整个的关闭过程
	stopWait sync.WaitGroup
	stopOnce sync.Once
	stopCh   chan struct{}
	stopped  atomic.Bool
}

func NewBatchProcessor() Processor {
	bsp := &batchDataProcessor{}

	bsp.stopWait.Add(1)
	go func() {
		defer bsp.stopWait.Done()
		bsp.processQueue()
		bsp.drainQueue()
	}()

	return bsp
}

func (bsp *batchDataProcessor) OnStart(ctx context.Context, data Data) {}

// Shutdown 优雅关闭处理器, 确保所有的 span 处理完成
func (bsp *batchDataProcessor) Shutdown(ctx context.Context) error {
	var err error
	// 如果存在多个线程, 那么其实最终只有一个线程是可以关闭的
	bsp.stopOnce.Do(func() {
		// 首先设置stopped 为 true
		bsp.stopped.Store(true)
		wait := make(chan struct{})
		go func() {
			//首先
			close(bsp.stopCh)
			bsp.stopWait.Wait()
			if bsp.e != nil {
			}
			close(wait)
		}()

		select {
		case <-wait:
		case <-ctx.Done():
			err = ctx.Err()
		}
	})

	return err
}

func (bsp *batchDataProcessor) ForceFlush(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if bsp.stopped.Load() {
		return nil
	}
	var err error
	if bsp.e != nil {
		flushCh := make(chan struct{})
	}

}

// OnEnd 将**结束的Span** 加入队列, 等待批量处理
func (bsp *batchDataProcessor) OnEnd(s Data) {
	//如果处理器已经关闭直接返回
	if bsp.stopped.Load() {
		return
	}

	if bsp.e == nil {
		return
	}

	bsp.enqueue(s)
}

/*
这段代码用处理Span数据的批量导出, 它的核心逻辑是通过事件驱动的办法
监听三个关键事件: 停止信号、定时器超时、队列数据到达

1. 如果是停止信号, 当这个通道关闭之后, 是可以立刻监听到的
2. 如果是定时器超时, 表示我们设置每次导出的时间
3. 队列数据到达, 应为当前的队列满了
*/
func (bsp *batchDataProcessor) processQueue() {
	defer bsp.timer.Stop()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		select {
		case <-bsp.stopCh:
			return
		case <-bsp.timer.C:
			if err := bsp.Export(ctx); err != nil {
				//todo
			}
		case sd := <-bsp.queue:
			bsp.batchMutex.Lock()
			bsp.batch = append(bsp.batch, sd)
			// 这里我们需要去判断当前是否应该导出
			shouldExport := len(bsp.batch) >= bsp.o.MaxExportBatchSize
			bsp.batchMutex.Unlock()
			// 如果应该导出
			if shouldExport {
				//判断我们的定时器是否停止
				if !bsp.timer.Stop() {
					select {
					case <-bsp.timer.C:
					default:
					}
				}
				if err := bsp.Export(ctx); err != nil {
					//todo:
				}
			}
		}
	}
}

// Export 导出对应的数据, 主要完成下面这几件事情
// 1. 重新去清除我们的定时器
// 2. 设置单次导出的的超时时间
// 3. 开始导出我们的span
func (bsp *batchDataProcessor) Export(ctx context.Context) error {
	//1) 首先我们应该清除对应的定时器
	bsp.timer.Reset(bsp.o.BatchTimeout)

	//然后 我们需要判断导出时间是否超时
	if bsp.o.ExportTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, bsp.o.ExportTimeout)
		defer cancel()
	}

	//2) 构造一个临界区
	bsp.batchMutex.Lock()
	defer bsp.batchMutex.Unlock()

	if l := len(bsp.batch); l > 0 {
		err := bsp.e.ExportSpans(ctx, bsp.batch)
		clear(bsp.batch)
		bsp.batch = bsp.batch[:0]
		if err != nil {
			return err
		}
	}
	return nil
}

/*
drainQueue 函数的目标是在停止入队后，清空队列中的剩余数据并导出最终的批次。
*/
func (bsp *batchDataProcessor) drainQueue() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		select {
		case sd := <-bsp.queue:
			bsp.batchMutex.Lock()
			bsp.batch = append(bsp.batch, sd)
			shouldExport := len(bsp.batch) == bsp.o.MaxExportBatchSize
			bsp.batchMutex.Unlock()
			if shouldExport {
				if err := bsp.Export(ctx); err != nil {
					//todo
				}
			}
		default:
			if err := bsp.Export(ctx); err != nil {
				//todo
			}
			return
		}
	}
}

// enqueue 主要的作用是将数据入队列
func (bsp *batchDataProcessor) enqueue(sd Data) {
	// 首先这里会创建对应的一个空的上下文
	// 之所以这样做的目的就是
	ctx := context.TODO()
	// 判断当前队列是阻塞队列还是非阻塞队列
	// 如果是阻塞队列那么就会等到队列可以插入数据的时候
	if bsp.o.BlockOnQueueFull {
		bsp.enqueueBlockOnQueueFull(ctx, sd)
	} else {
		//如果是非阻塞队列, 那么就会
		bsp.enqueueDrop(ctx, sd)
	}
}

/*
这里其实是存在两种行为的,
1. 第一种就是成功写入
2. 第二种就是直接丢弃, 然后增加对应的计数
*/
func (bsp *batchDataProcessor) enqueueDrop(_ context.Context, sd Data) bool {
	select {
	case bsp.queue <- sd:
		return true
	default:
		atomic.AddUint32(&bsp.dropped, 1)
	}
	return false
}

/*
这种事阻塞队列
如果当前的队列一直无法写入
*/
func (bsp *batchDataProcessor) enqueueBlockOnQueueFull(ctx context.Context, data Data) bool {
	select {
	case bsp.queue <- data:
		return true
	case <-ctx.Done():
		return false
	}
}
