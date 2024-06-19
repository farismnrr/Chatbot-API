package smtp

import (
	"capstone-project/helper"
	"io/ioutil"
	"log"
	"net/smtp"
	"strconv"
	"strings"
)

func SendMailSimple(subject string, otp int, to string) {
	smtpHost := helper.GetEnv("SMTP_HOST")
	smtpPassword := helper.GetEnv("SMTP_PASSWORD")
	smtpUser := helper.GetEnv("SMTP_USER")

	html, err := ioutil.ReadFile("smtp/smtp.html")
	if err != nil {
		log.Fatal(err)
	}

	// Replace the OTP value in the HTML content
	newOtpValue := strconv.Itoa(otp)
	htmlStr := string(html)
	htmlStr = strings.ReplaceAll(htmlStr, "{{OTP}}", newOtpValue)

	header := "MIME-Version: 1.0\r\n"
	header += "Content-Type: text/html; charset=utf-8\r\n"
	header += "Subject: " + subject + "\r\n"
	msg := header + "\r\n" + htmlStr

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	err = smtp.SendMail(smtpHost+":587", auth, smtpUser, []string{to}, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}
