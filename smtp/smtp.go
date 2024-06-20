package smtp

import (
	"capstone-project/helper"
	"io/ioutil"
	"net/smtp"
	"strings"
)

func SendMailSimple(subject string, otp string, to string) error {
	smtpHost := helper.GetEnv("SMTP_HOST")
	smtpPassword := helper.GetEnv("SMTP_PASSWORD")
	smtpUser := helper.GetEnv("SMTP_USER")
	smtpPort := helper.GetEnv("SMTP_PORT")

	html, err := ioutil.ReadFile("smtp/smtp.html")
	if err != nil {
		return err
	}

	// Replace the OTP value in the HTML content
	htmlStr := string(html)
	htmlStr = strings.ReplaceAll(htmlStr, "{{OTP}}", otp)

	header := "MIME-Version: 1.0\r\n"
	header += "Content-Type: text/html; charset=utf-8\r\n"
	header += "Subject: " + subject + "\r\n"
	msg := header + "\r\n" + htmlStr

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
