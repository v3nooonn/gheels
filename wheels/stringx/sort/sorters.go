package sort

import (
	"strings"
)

var (
	order = "abcdefghijklmnopqrstuvwxyzæøå"
)

type defaultSorter struct {
	order string
}

func (s defaultSorter) Sort(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}

	s1Rune := []rune(strings.ToLower(s1))
	s2Rune := []rune(strings.ToLower(s2))

	s1Len := len(s1Rune)
	s2Len := len(s2Rune)
	min := 0
	if s1Len >= s2Len {
		min = s2Len
	} else {
		min = s1Len
	}

	for i := 0; i < min; i++ {
		if s1Rune[i] == s2Rune[i] {
			continue
		}

		s1Idx := strings.Index(s.order, string(s1Rune[i]))
		s2Idx := strings.Index(s.order, string(s2Rune[i]))
		if s1Idx > 0 && s2Idx > 0 {
			return s1Idx <= s2Idx
		}

		return s1Rune[i] <= s2Rune[i]
	}

	return s1Len < s2Len
}
