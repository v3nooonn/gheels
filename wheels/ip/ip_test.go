package ip

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	one = "11.11.11.11"
	two = "22.22.22.22"
)

func FuzzIPParseCase1(f *testing.F) {
	f.Add("fake")
	f.Fuzz(func(t *testing.T, s1 string) {
		ipAddress, _ := Parse(&http.Request{
			Header: map[string][]string{
				"X-Forwarded-For": {one, two},
			},
			RemoteAddr: "127.0.0.1:123",
		})
		assert.Equal(t, one, ipAddress)
	})
}

func FuzzIPParseCase2(f *testing.F) {
	f.Add("fake")
	f.Fuzz(func(t *testing.T, s1 string) {
		ipAddress, _ := Parse(&http.Request{
			Header: map[string][]string{
				"X-Forwarded-For": {s1, one, two},
			},
			RemoteAddr: "127.0.0.1:123",
		})
		assert.Equal(t, local, ipAddress)
	})
}
func FuzzIPParseFailedCase1(f *testing.F) {
	f.Add("fake")
	f.Fuzz(func(t *testing.T, s1 string) {
		ipAddress, err := Parse(&http.Request{
			Header: map[string][]string{
				"X-Forwarded-For": {s1, one, two},
			},
			RemoteAddr: s1,
		})
		assert.Equal(t, "", ipAddress)
		assert.EqualError(t, fmt.Errorf("parsing remote address occurred an error: address fake: missing port in address"), err.Error())
	})
}
