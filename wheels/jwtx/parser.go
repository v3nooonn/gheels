package jwtx

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
)

func init() {
	ES256 = &es256Parser{
		Header: Header{
			Algorithm: AlgEs256,
			Type:      "JWT",
		},
	}
}

type (
	Parser interface {
		Parse(s string) (string, bool)
	}

	es256Parser struct {
		Header Header
	}

	Header struct {
		Algorithm Algorithm `json:"alg"`
		Type      string    `json:"typ"`
	}

	Algorithm string
)

const (
	AlgEs256 Algorithm = "ES256"
	AlgHs256 Algorithm = "HS256"
)

var ES256 *es256Parser

// Parse parsing and basic validating token string from Authentication header
func (p *es256Parser) Parse(s string) (string, bool) {
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
