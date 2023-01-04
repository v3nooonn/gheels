package jwtx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzJwtParse(f *testing.F) {
	f.Add("Bearer ey.in.ci")
	f.Fuzz(func(t *testing.T, jwt string) {
		_, ok := Parse(jwt)
		assert.Equal(t, true, ok, fmt.Sprintf("input: %s, result: %v\n", jwt, ok))
	})
}

func FuzzJwtParseFailedCase1(f *testing.F) {
	f.Add("freestyle")
	f.Fuzz(func(t *testing.T, jwt string) {
		_, ok := Parse(jwt)
		assert.Equal(t, false, ok, fmt.Sprintf("input: %s, result: %v\n", jwt, ok))
	})
}

func FuzzJwtParseFailedCase2(f *testing.F) {
	f.Add("Bearer")
	f.Fuzz(func(t *testing.T, jwt string) {
		_, ok := Parse(jwt)
		assert.Equal(t, false, ok, fmt.Sprintf("input: %s, result: %v\n", jwt, ok))
	})
}

func FuzzJwtParseFailedCase3(f *testing.F) {
	f.Add("Bearer ")
	f.Fuzz(func(t *testing.T, jwt string) {
		_, ok := Parse(jwt)
		assert.Equal(t, false, ok, fmt.Sprintf("input: %s, result: %v\n", jwt, ok))
	})
}

func FuzzJwtParseFailedCase4(f *testing.F) {
	f.Add("Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9")
	f.Fuzz(func(t *testing.T, jwt string) {
		_, ok := Parse(jwt)
		assert.Equal(t, false, ok, fmt.Sprintf("input: %s, result: %v\n", jwt, ok))
	})
}

func FuzzJwtParseFailedCase5(f *testing.F) {
	f.Add("Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.")
	f.Fuzz(func(t *testing.T, jwt string) {
		_, ok := Parse(jwt)
		assert.Equal(t, false, ok, fmt.Sprintf("input: %s, result: %v\n", jwt, ok))
	})
}

func FuzzJwtParseFailedCase6(f *testing.F) {
	f.Add("Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.x")
	f.Fuzz(func(t *testing.T, jwt string) {
		_, ok := Parse(jwt)
		assert.Equal(t, false, ok, fmt.Sprintf("input: %s, result: %v\n", jwt, ok))
	})
}

func FuzzJwtParseFailedCase7(f *testing.F) {
	f.Add("Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.x.")
	f.Fuzz(func(t *testing.T, jwt string) {
		_, ok := Parse(jwt)
		assert.Equal(t, false, ok, fmt.Sprintf("input: %s, result: %v\n", jwt, ok))
	})
}
