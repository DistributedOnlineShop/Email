package gapi

import (
	"fmt"
	"net/smtp"
)

func (s *Server) EmailHeader(context EmailContext) error {
	auth := smtp.PlainAuth("", s.config.SmtpUser, s.config.SmtpPassword, s.config.SmtpHost)

	message := "Subject: " + context.Subject + "\n" + "Content-Type: text/html; charset=utf-8" + "\n" + context.Body

	err := smtp.SendMail(s.config.SmtpHost+":"+s.config.SmtpPort, auth, s.config.SmtpUser, context.Email, []byte(message))
	if err != nil {
		return fmt.Errorf("send email failed")
	}

	return nil
}
