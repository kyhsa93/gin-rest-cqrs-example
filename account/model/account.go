package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Account account model
type Account struct {
	ID          string    `json:"id" example:"accountId"`
	ProfileID   []string  `json:"profileId" example:"profileId"`
	Email       string    `json:"email" example:"test@gmail.com"`
	Provider    string    `json:"provider" exmaple:"gmail"`
	AccessToken string    `json:"accessToken" example:"accesstoken"`
	CreatedAt   time.Time `json:"createdAt" example:"2019-12-23 12:27:37"`
	UpdatedAt   time.Time `json:"updatedAt" example:"2019-12-23 12:27:37"`
}

// CreateAccessToken create access token with jwt
func (account *Account) CreateAccessToken() string {
	expirationTime := time.Now().Add(500 * time.Minute)
	claims := jwt.StandardClaims{ExpiresAt: expirationTime.Unix(), Issuer: account.ID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, tokenError := token.SignedString([]byte("access token secret"))

	if tokenError != nil {
		panic(tokenError)
	}

	return tokenString
}

// ValidateAccessToken validation token in Account model
func (account *Account) ValidateAccessToken() bool {
	if account.AccessToken == "" {
		return false
	}
	claims := &jwt.StandardClaims{}
	jwtToken, _ := jwt.ParseWithClaims(
		account.AccessToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte("access token secret"), nil
		})
	if jwtToken.Valid == true && claims.Issuer != "" {
		return true
	}
	return false
}
