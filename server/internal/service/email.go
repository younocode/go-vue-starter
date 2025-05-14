package service

import (
	"fmt"
	"github.com/younocode/go-vue-starter/server/config"
	"net/smtp"

	mail "github.com/jordan-wright/email"
)

type EmailSend struct {
	addr    string
	myEmail string
	subject string
	auth    smtp.Auth
}

func NewEmailSend(cfg *config.EmailConfig) *EmailSend {
	return &EmailSend{
		addr: fmt.Sprintf(""),
	}
}

func (s *EmailSend) Send(email string, emailCode string) error {
	instance := mail.NewEmail()
	instance.From = s.myEmail
	instance.To = []string{email}
	instance.Subject = s.subject
	instance.Text = []byte(fmt.Sprintf("你的验证码： %s!", emailCode))
	return instance.Send(s.addr, s.auth)
}
