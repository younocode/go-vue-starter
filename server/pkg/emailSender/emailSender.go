package emailSender

import (
	"fmt"
	mail "github.com/jordan-wright/email"
	"github.com/younocode/go-vue-starter/server/config"
	"net/smtp"
)

type EmailSender struct {
	addr    string
	myMail  string
	subject string
	auth    smtp.Auth
}

func NewEmailSend(cfg config.EmailConfig) *EmailSender {
	return &EmailSender{
		addr:    cfg.DSN(),
		auth:    smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host),
		myMail:  cfg.Username,
		subject: cfg.Subject,
	}
}

func (e *EmailSender) Send(email string, emailCode string) error {
	instance := mail.NewEmail()
	instance.From = e.myMail
	instance.To = []string{email}
	instance.Subject = e.subject
	instance.Text = []byte(fmt.Sprintf("你的验证码为： %s", emailCode))

	return instance.Send(e.addr, e.auth)
}
