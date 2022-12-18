package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("ashdjqy9283409bsdklkg8hda01sss")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
