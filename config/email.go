package config

import (
	"github.com/caarlos0/env"
)

// Email email configuration struct
type Email struct {
	Address  string `env:"EMAIL_ADDRESS" envDefault:""`
	Password string `env:"EMAIL_PASSWORD" envDefault:""`
	SMTPHost string `env:"SMTP_HOST" envDefault:"smtp.gmail.com"`
	SMTPPort string `env:"SMTP_PORT" envDefault:"587"`
}

// NewEmail create email config instance
func NewEmail() *Email {
	email := &Email{}
	env.Parse(email)
	return email
}
