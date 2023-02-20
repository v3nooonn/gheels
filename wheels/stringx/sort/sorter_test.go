package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzStrSorter(f *testing.F) {
	f.Add("asdfadsf", "zxc")
	f.Fuzz(func(t *testing.T, s1, s2 string) {
		result := NewSorter(CustomizedSorter).Sort(s1, s2)
		assert.Equal(t, true, result, fmt.Sprintf("input: %s, result: %v\n", s1, s2))
	})
}
