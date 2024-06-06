package internal

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	UserID          uint   `json:"user_id"`
	Username        string `json:"username"`
	UserRole        string `json:"user_role"`
	TokenBufferTime int64  `json:"token_buffer_time"`
	jwt.StandardClaims
}
