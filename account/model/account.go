package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Account account model
type Account struct {
	ID        string    `json:"id" example:"389df385-ccaa-49c1-aee2-698ba1191857"`
	Email     string    `json:"email" example:"test@test.com"`
	CreatedAt time.Time `json:"created_at" example:"2019-12-23 12:27:37"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-12-23 12:27:37"`
}

// CreateAccessToken create access token with jwt
func (account *Account) CreateAccessToken() string {
	claims := jwt.StandardClaims{ExpiresAt: 15000, Issuer: account.ID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, tokenError := token.SignedString([]byte("access token secret"))

	if tokenError != nil {
		panic(tokenError)
	}

	return tokenString
}

// CreateRefreshToken create refresh token with jwt
func (account *Account) CreateRefreshToken(accessToken string) string {
	claims := jwt.StandardClaims{ExpiresAt: 150000, Issuer: accessToken}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, tokenError := token.SignedString([]byte("refresh token secret"))

	if tokenError != nil {
		panic(tokenError)
	}

	return tokenString
}
