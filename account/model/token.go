package model

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

// Token access and refresh
type Token struct {
	ID     string
	Access string
}

// Validate validate token
func (token *Token) Validate() string {
	if token.ID == "" || token.Access == "" {
		return ""
	}
	claims := &jwt.StandardClaims{}
	jwtToken, _ := jwt.ParseWithClaims(token.Access, claims, func(accessToken *jwt.Token) (interface{}, error) {
		log.Println(token)
		return []byte("access token secret"), nil
	})

	if jwtToken.Valid == true && claims.Issuer != "" {
		return claims.Issuer
	}

	return ""
}
