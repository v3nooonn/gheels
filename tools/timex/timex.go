package timex

import (
	"math"
	"time"

	"github.com/v3nooonn/gheels/tools/errorx"
)

type CalendarUnit int

const (
	Unknown CalendarUnit = iota
	Year
	Month
	Day
)

// DiffAbsByFloor
// 1. Both of `FROM` and `TO` should in a same timezone.
// If the timezone is customized, they should be formed with the same name and offset.
// 2. Supported units contains: year, month and day only.
func DiffAbsByFloor(from, to time.Time, unit CalendarUnit) (int, error) {
	if from.Location().String() != to.Location().String() {
		return 0, errorx.Internal("timezone mismatch")
	}

	if from.Equal(to) {
		return 0, nil
	}
	if from.After(to) {
		from, to = to, from
	}

	switch unit {
	case Year:
		return byYearFloor(from, to), nil
	case Month:
		return byMonthFloor(from, to), nil
	case Day:
		return byDayFloor(from, to), nil
	default:
		return 0, errorx.Internal("unsupported unit")
	}
}

func byYearFloor(from, to time.Time) int {
	years := to.Year() - from.Year()
	if years == 0 {
		return 0
	}

	diffAdded := from.AddDate(years, 0, 0)
	if diffAdded.After(to) {
		return years - 1
	}

	return years
}

func byMonthFloor(from, to time.Time) int {
	years := byYearFloor(from, to)
	if years > 0 {
		from = from.AddDate(years, 0, 0)
	}

	months := 0
	for {
		from = from.AddDate(0, 1, 0)
		if from.After(to) {
			break
		}
		months++
	}

	return (12 * years) + months
}

func byDayFloor(from, to time.Time) int {
	return int(math.Floor(float64(to.Sub(from) / (24 * time.Hour))))
}
