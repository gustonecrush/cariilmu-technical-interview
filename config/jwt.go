package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("afafjija9234nkaf")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}