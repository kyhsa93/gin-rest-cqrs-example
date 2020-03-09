package config

import (
	"os"
	"strconv"
)

// AuthConfigInterface auth config interface
type AuthConfigInterface interface {
	AccessTokenSecret() string
	AccessTokenExpiration() int
}

// Auth auth config struct
type Auth struct {
	accessTokenSecret     string `env:"ACCESS_SECRET" envDefault:"ACCESS_SECRET"`
	accessTokenExpiration string `env:"ACCESS_EXPIRATION" envDefault:"60"`
}

// NewAuthConfig create Auth cconfig instance
func NewAuthConfig() *Auth {
	accessTokenSecret := "accessTokenSecret"
	accessTokenExpiration := "60"

	if env := os.Getenv("ACCESS_TOKEN_SECRET"); env != "" {
		accessTokenSecret = env
	}
	if env := os.Getenv("ACCESS_TOKEN_ENPIRATION"); env != "" {
		accessTokenExpiration = env
	}
	auth := &Auth{
		accessTokenSecret:     accessTokenSecret,
		accessTokenExpiration: accessTokenExpiration,
	}
	return auth
}

//AccessTokenSecret get accessToken secret
func (auth *Auth) AccessTokenSecret() string {
	return auth.accessTokenSecret
}

//AccessTokenExpiration get accessToken expiration
func (auth *Auth) AccessTokenExpiration() int {
	minutes, err := strconv.Atoi(auth.accessTokenExpiration)
	if err != nil {
		panic(err)
	}
	return minutes
}
