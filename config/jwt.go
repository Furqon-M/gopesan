package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("ashdjqy9283409bsdk1kg8da01")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
