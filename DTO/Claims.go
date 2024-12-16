package DTO

import "github.com/golang-jwt/jwt/v5"

type DicodeClaims struct {
	Username  string   `json:"username"`
	Roles     []string `json:"roles"`
	jwt.RegisteredClaims
}
