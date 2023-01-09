package timex

import (
	"fmt"
	"testing"
	"time"

	"github.com/v3nooonn/gheels/tools/constant"

	"github.com/stretchr/testify/assert"
)

func FuzzDiffByYearsCase1(f *testing.F) {
	f.Add("2000-01-01", "2001-02-02")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Year)
		assert.Equal(t, 1, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByYearsCase2(f *testing.F) {
	f.Add("2003-01-01", "2001-02-02")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Year)
		assert.Equal(t, 1, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByYearsCase3(f *testing.F) {
	f.Add("2003-03-01", "2001-02-02")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Year)
		assert.Equal(t, 2, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByYearsCase4(f *testing.F) {
	f.Add("2000-01-01", "2001-01-01")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Year)
		assert.Equal(t, 1, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByMonthsCase1(f *testing.F) {
	f.Add("2000-01-01", "2001-01-01")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Month)
		assert.Equal(t, 12, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByMonthsCase2(f *testing.F) {
	f.Add("2000-01-02", "2001-01-01")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Month)
		assert.Equal(t, 11, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByMonthsCase3(f *testing.F) {
	f.Add("2001-01-01", "2000-01-10")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Month)
		assert.Equal(t, 11, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByMonthsCase4(f *testing.F) {
	f.Add("2001-01-10", "2001-10-01")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Month)
		assert.Equal(t, 8, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByMonthsCase5(f *testing.F) {
	f.Add("2001-01-10", "2001-10-11")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Month)
		assert.Equal(t, 9, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByMonthsCase6(f *testing.F) {
	f.Add("2000-01-10", "2001-10-11")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Month)
		assert.Equal(t, 21, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByMonthsCase7(f *testing.F) {
	f.Add("2001-10-11", "2000-01-10")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Month)
		assert.Equal(t, 21, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByDaysCase1(f *testing.F) {
	f.Add("2000-01-10", "2000-01-10")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Day)
		assert.Equal(t, 0, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByDaysCase2(f *testing.F) {
	f.Add("2000-01-10", "2000-01-11")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Day)
		assert.Equal(t, 1, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByDaysCase3(f *testing.F) {
	f.Add("2000-01-12", "2000-01-11")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMD.String(), fs)
		to, _ := time.Parse(constant.YMD.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Day)
		assert.Equal(t, 1, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByDaysCase4(f *testing.F) {
	f.Add("2000-01-11 11:00:00", "2000-01-12 12:00:00")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMDHMS.String(), fs)
		to, _ := time.Parse(constant.YMDHMS.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Day)
		assert.Equal(t, 1, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}

func FuzzDiffByDaysCase5(f *testing.F) {
	f.Add("2000-01-11 13:00:00", "2000-01-12 12:00:00")
	f.Fuzz(func(t *testing.T, fs, ts string) {
		from, _ := time.Parse(constant.YMDHMS.String(), fs)
		to, _ := time.Parse(constant.YMDHMS.String(), ts)

		diffs, _ := DiffAbsByFloor(from, to, Day)
		assert.Equal(t, 0, diffs,
			fmt.Sprintf("input: %s, %s, result: %v\n", fs, ts, diffs))
	})
}
