package email

import (
	"log"
	"net/smtp"

	"github.com/caarlos0/env"
)

// Email email configuration struct
type Email struct {
	Address  string `env:"EMAIL_ADDRESS" envDefault:""`
	Password string `env:"EMAIL_PASSWORD" envDefault:""`
	SMTPHost string `env:"SMTP_HOST" envDefault:"smtp.gmail.com"`
	SMTPPort string `env:"SMTP_PORT" envDefault:"587"`
}

// Send send email with template
func (email *Email) Send(receiver []string) {
	if email.Address == "" {
		log.Println("Can not use SMTP server")
		return
	}

	emailAuth := smtp.PlainAuth(
		"",
		email.Address,
		email.Password,
		email.SMTPHost,
	)

	message := []byte("test email message")

	err := smtp.SendMail(
		email.SMTPHost+":"+email.SMTPPort,
		emailAuth,
		email.Address,
		receiver,
		message,
	)

	if err != nil {
		log.Println(err)
	}
}

// New create email config instance
func New() *Email {
	email := &Email{}
	env.Parse(email)
	return email
}
