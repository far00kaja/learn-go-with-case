package dto

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Tokens struct {
	Authorization string `header:"Authorization" binding:"required"`
}

type TokensResponse struct {
	Token string    `json:"token"`
	ID    uuid.UUID `json:"id"`
	jwt.StandardClaims
}
