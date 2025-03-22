package model

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	jwt.RegisteredClaims
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
