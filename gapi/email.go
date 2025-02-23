package gapi

import (
	"fmt"
	"net/smtp"

	"Email/util"
)

func SendEmail(config util.Config, context EmailContext) error {
	auth := smtp.PlainAuth("", config.SmtpUser, config.SmtpPassword, config.SmtpHost)

	message := "Subject: " + context.Subject + "\n" + "Content-Type: text/html; charset=utf-8" + "\n" + context.Body

	err := smtp.SendMail(config.SmtpHost+":"+config.SmtpPort, auth, config.SmtpUser, context.To, []byte(message))
	if err != nil {
		return fmt.Errorf("send email failed")
	}

	return nil
}
