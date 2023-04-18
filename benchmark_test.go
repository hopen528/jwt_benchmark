package test

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func BenchmarkSigned(b *testing.B) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	for n := 0; n < b.N; n++ {
		sigOptions := (&jose.SignerOptions{}).WithType("JWT").WithHeader("kid", "random uuid")
		sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.PS256, Key: privateKey}, sigOptions)
		if err != nil {
			panic(err)
		}

		cl := jwt.Claims{
			Subject:   "subject",
			Issuer:    "issuer",
			NotBefore: jwt.NewNumericDate(time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)),
			Audience:  jwt.Audience{"leela", "fry"},
		}
		_, err = jwt.Signed(sig).Claims(cl).CompactSerialize()

		if err != nil {
			panic(err)
		}
	}
}
