package mailer

import (
	"context"

	"github.com/kzmijak/zswod_api_go/modules/errors"
	"gopkg.in/gomail.v2"
)

const (
	ErrFailedToSendMessage = "ErrFailedToSendMessage: Failed to send the provided message"
)

type Mailer struct {
	dialer *gomail.Dialer
}

func Initialize(ctx context.Context, cfg MailerConfig) *Mailer {
	dialer := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	
	return &Mailer{
		dialer: dialer,
	}
}

type MailerMessage struct {
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	Cc string `json:"cc"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}
func (m Mailer) Send(r MailerMessage) (ok bool, err error) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", r.Sender)
	msg.SetHeader("To", r.Receiver)
	msg.SetHeader("Cc", r.Cc)
	msg.SetHeader("Subject", r.Subject)
	msg.SetHeader("text/html", r.Content)

	if err := m.dialer.DialAndSend(msg); err != nil {
		return false, errors.Error(ErrFailedToSendMessage)
	}

	return true, nil
}