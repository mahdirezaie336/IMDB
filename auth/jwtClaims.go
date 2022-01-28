package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserId   int    `json:"user_id"`
	jwt.StandardClaims
}
