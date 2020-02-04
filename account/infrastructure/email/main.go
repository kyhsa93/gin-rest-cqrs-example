package email

import (
	"log"
	"net/smtp"

	"github.com/kyhsa93/gin-rest-example/config"
)

// Email email struct
type Email struct {
	Address  string
	Password string
	SMTPHost string
	SMTPPort string
}

// New create email instance
func New(config *config.Config) *Email {
	return &Email{
		Address:  config.Email.Address,
		Password: config.Email.Password,
		SMTPHost: config.Email.SMTPHost,
		SMTPPort: config.Email.SMTPPort,
	}
}

// Send email message to receivers
func (email *Email) Send(receivers []string, message string) {
	if email.Address == "" {
		log.Println("Email Credential is missing.")
		return
	}

	emailAuth := smtp.PlainAuth(
		"",
		email.Address,
		email.Password,
		email.SMTPHost,
	)

	err := smtp.SendMail(
		email.SMTPHost+":"+email.SMTPPort,
		emailAuth,
		email.Address,
		receivers,
		[]byte(message),
	)

	if err != nil {
		log.Println(err)
	}
}
