package mailer

import (
	"crypto/tls"

	"github.com/asma12a/challenge-s6/config"
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	d := gomail.NewDialer("smtp-relay.brevo.com", 587, "bastiendikiadi@gmail.com", config.Env.BrevoAPIKey)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "squadgo@squadgo.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
