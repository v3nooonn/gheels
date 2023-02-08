package jwtx

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// Assigner jwt token assigner
type Assigner interface {
	Assign() (string, error)
}

type CustomizedAssigner struct {
	jwt.StandardClaims
	Key string
}

type CustomizedClaim struct {
	jwt.StandardClaims
	// fields...
}

func (c *CustomizedAssigner) Assign() (string, error) {
	rsaKey, err := jwt.ParseECPrivateKeyFromPEM([]byte("PrivateKey"))
	if err != nil {
		return "", errors.Wrap(err, err.Error())
	}

	now := time.Now()
	claim := CustomizedClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:   "r.Issuer",
			Subject:  "r.Subject",
			Audience: "r.Audience",
			IssuedAt: now.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(jwt.SigningMethodES256.Name), claim)

	return token.SignedString(rsaKey)
}
