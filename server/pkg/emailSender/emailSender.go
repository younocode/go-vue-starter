package emailSender

import (
	"fmt"
	mail "github.com/jordan-wright/email"
	"github.com/younocode/go-vue-starter/server/config"
	"net/smtp"
	"strings"
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

	err := instance.Send(e.addr, e.auth)
	// qq邮箱 在 QUIT 时返回 “[]”，引起 short response: [] 错误
	// net/textproto/reader.go/parseCodeLine 部分
	if strings.Contains(err.Error(), "short response") {
		return nil
	}
	return err
}
