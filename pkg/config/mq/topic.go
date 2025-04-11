package mq

type Topic struct {
	Name              string
	NumPartition      int32
	ReplicationFactor int16
}

type TopicOption interface {
	Option(topic Topic) Topic
}

type OptionFunc func(topic Topic) Topic

func (f OptionFunc) Option(topic Topic) {
	f(topic)
}

func WithNumPartition(NumPartition int32) OptionFunc {
	return OptionFunc(func(topic Topic) Topic {
		topic.NumPartition = NumPartition
		return topic
	})
}

func WithReplicationFactor(ReplicationFactor int16) OptionFunc {
	return OptionFunc(func(topic Topic) Topic {
		topic.ReplicationFactor = ReplicationFactor
		return topic
	})
}

func NewTopic(name string, opts ...TopicOption) Topic {
	topic := Topic{
		Name:              name,
		NumPartition:      1,
		ReplicationFactor: 1,
	}

	for _, o := range opts {
		topic = o.Option(topic)
	}

	return topic
}
