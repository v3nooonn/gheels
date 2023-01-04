package constant

type TimeFormat string

const (
	MD     TimeFormat = "01-02"
	HMS    TimeFormat = "15:04:05"
	YM     TimeFormat = "2006-01"
	YMD    TimeFormat = "2006-01-02"
	YMDHM  TimeFormat = "2006-01-02 15:00"
	YMDHMZ TimeFormat = "2006-01-02 15:00 -07:00"
	YMDHMS TimeFormat = "2006-01-02 15:04:05"
)

func (u TimeFormat) String() string {
	return string(u)
}
