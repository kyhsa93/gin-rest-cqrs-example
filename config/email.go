package config

import "os"

// EmailConfigInterface email config interface
type EmailConfigInterface interface {
	Address() string
	Password() string
	SMTPHost() string
	SMTPPort() string
}

// Email email configuration struct
type Email struct {
	address  string `env:"EMAIL_ADDRESS" envDefault:""`
	password string `env:"EMAIL_PASSWORD" envDefault:""`
	smtpHost string `env:"SMTP_HOST" envDefault:"smtp.gmail.com"`
	smtpPort string `env:"SMTP_PORT" envDefault:"587"`
}

// NewEmailConfig create email config instance
func NewEmailConfig() *Email {
	address := ""
	password := ""
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	if env := os.Getenv(""); env != "" {
		address = env
	}
	if env := os.Getenv(""); env != "" {
		password = env
	}
	if env := os.Getenv(""); env != "" {
		smtpHost = env
	}
	if env := os.Getenv(""); env != "" {
		smtpPort = env
	}

	email := &Email{
		address:  address,
		password: password,
		smtpHost: smtpHost,
		smtpPort: smtpPort,
	}
	return email
}

// Address get email Address
func (email *Email) Address() string {
	return email.address
}

// Password get email Password
func (email *Email) Password() string {
	return email.password
}

// SMTPHost get email SMTPHost
func (email *Email) SMTPHost() string {
	return email.smtpHost
}

// SMTPPort get email SMTPPort
func (email *Email) SMTPPort() string {
	return email.smtpPort
}
