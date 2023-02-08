package jwtx

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
)

type (
	Parser interface {
		Parse(s string) (string, bool)
	}

	ES256Parser struct {
		Header JWTHeader
	}

	JWTHeader struct {
		Algorithm Algorithm `json:"alg"`
		Type      string    `json:"typ"`
	}

	Algorithm string
)

const (
	ES256 Algorithm = "ES256"
	HS256 Algorithm = "HS256"
)

func NewES256Parser() Parser {
	return &ES256Parser{
		Header: JWTHeader{
			Algorithm: ES256,
			Type:      "JWT",
		},
	}
}

// Parse parsing and basic validating token string from Authentication header
func (p *ES256Parser) Parse(s string) (string, bool) {
	bytes, err := json.Marshal(p.Header)
	if err != nil {
		return "", false
	}
	algHeader := base64.StdEncoding.EncodeToString(bytes)

	regFmt := `^(Bearer ){1}((%s)+)(\.){1}([-_a-zA-Z0-9]+)(\.){1}([-_a-zA-Z0-9]+)$`

	jwtPrefix := regexp.MustCompile(`^Bearer `)
	jwtPattern := regexp.MustCompile(fmt.Sprintf(regFmt, algHeader))

	return string(jwtPrefix.ReplaceAll([]byte(s), []byte(""))), jwtPattern.Match([]byte(s))
}
