package constant

type TimeLayout string

const (
	MD     TimeLayout = "01-02"
	HMS    TimeLayout = "15:04:05"
	YM     TimeLayout = "2006-01"
	YMD    TimeLayout = "2006-01-02"
	YMDHM  TimeLayout = "2006-01-02 15:00"
	YMDHMZ TimeLayout = "2006-01-02 15:00 -07:00"
	YMDHMS TimeLayout = "2006-01-02 15:04:05"
)

func (u TimeLayout) String() string {
	return string(u)
}
