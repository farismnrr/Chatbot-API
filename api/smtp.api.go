package api

import (
	"capstone-project/helper"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strings"
)

func SendMailSimple(subject string, otp string, to string) error {
	smtpHost := helper.GetEnv("SMTP_HOST")
	smtpPassword := helper.GetEnv("SMTP_PASSWORD")
	smtpUser := helper.GetEnv("SMTP_USER")
	smtpPort := helper.GetEnv("SMTP_PORT")

	url := helper.GetEnv("SMTP_TEMPLATE_URL")
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	htmlStr := string(body)
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
