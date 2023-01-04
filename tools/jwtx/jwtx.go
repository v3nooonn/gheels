package jwtx

import "regexp"

func Parse(s string) (string, bool) {
	jwtPrefix := regexp.MustCompile(`^Bearer `)
	jwtPattern := regexp.MustCompile(`^(Bearer ){1}([-_a-zA-Z0-9]+)(\.){1}([-_a-zA-Z0-9]+)(\.){1}([-_a-zA-Z0-9]+)$$`)

	return string(jwtPrefix.ReplaceAll([]byte(s), []byte(""))), jwtPattern.Match([]byte(s))
}
