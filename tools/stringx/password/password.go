package password

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/v3nooonn/gheels/tools/errorx"
)

func ValidatePassword(pwd string, min, max int, symbols string) error {
	// Caused by the built-in regexp error: invalid or unsupported Perl syntax.
	// Below, is the regular expression for FE, this validation is based on it.
	// ^(?=.+[a-z])(?=.+[A-Z])(?=.+[0-9])(?=.+[~!@#$%^&*()_\+\-\= ,.\/<>?\[\]{}|;':"])[\da-zA-Z~!@#$%^&*()_\+\-\= ,.\/<>?\[\]{}|;':"]{12,16}$
	// The allowed symbols: ~!@#$%^&*()_+-= ,./<>?[]{}|;':"
	upper, lower, number, symbol := false, false, false, false
	for _, s := range pwd {
		switch {
		case unicode.IsUpper(s):
			upper = true
		case unicode.IsLower(s):
			lower = true
		case unicode.IsNumber(s):
			number = true
		default:
			if strings.Index(symbols, string(s)) <= 0 {
				return errorx.BadRequest(fmt.Sprintf("the password contains invalid symbol: %s", string(s)))
			}
			symbol = true
		}
	}
	if len(pwd) < min || len(pwd) > max {
		return errorx.BadRequest(fmt.Sprintf("the length of password should be in the range of %v to %v characters", min, max))
	}

	if !upper || !lower || !number || !symbol {
		return errorx.BadRequest("password should contains a upper letter, a lower letter, a number and a symbol at least")
	}

	return nil
}
