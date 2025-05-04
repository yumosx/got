package timex

import "time"

const (
	StandardDateFormat     = "2006-01-02"
	StandardDateTimeFormat = "2006-01-02 15:04:05"
)

// FormatDate 将时间格式化为标准日期字符串
func FormatDate(t time.Time) string {
	return t.Format(StandardDateFormat)
}

// ParseDate 解析日期字符串为时间对象
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse(StandardDateFormat, dateStr) // 修正了参数顺序错误
}

// FormatDateTime 将时间格式化为标准日期时间字符串
func FormatDateTime(t time.Time) string {
	return t.Format(StandardDateTimeFormat)
}

// UnixToDateString 将Unix时间戳转换为日期字符串
func UnixToDateString(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(StandardDateFormat)
}

// UnixToDateTimeString 将Unix时间戳转换为日期时间字符串
func UnixToDateTimeString(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(StandardDateTimeFormat)
}

// DateStringToUnix 将标准日期字符串转换成为unix 时间戳
func DateStringToUnix(dataStr string) (int64, error) {
	t, err := ParseDate(dataStr)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func FormatDateRange(start, end string) (time.Time, time.Time, error) {
	var (
		err error
		s   time.Time
		e   time.Time
	)

	s, err = ParseDate(start)
	e, err = ParseDate(end)

	return s, e, err
}
