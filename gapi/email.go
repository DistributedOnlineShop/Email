package api

import (
	"fmt"
	"net/smtp"

	"Email/util"
)

func SendEmail(config util.Config, from, body, subject string, to []string) error {
	auth := smtp.PlainAuth("", config.SmtpUser, config.SmtpPassword, config.SmtpHost)

	message := "Subject: " + subject + "\n" + "Content-Type: text/html; charset=utf-8" + body

	err := smtp.SendMail(config.SmtpHost+":"+config.SmtpPort, auth, from, to, []byte(message))
	if err != nil {
		return fmt.Errorf("send email failed")
	}

	return nil
}
