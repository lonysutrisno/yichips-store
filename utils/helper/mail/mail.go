package mail

import (
	// native

	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func SendMail(toId, subject, body string) {
	mail := Mail{}
	mail.senderId = os.Getenv("MAIL_ADDRESS")
	mail.toIds = []string{toId}
	mail.subject = subject
	mail.body = body

	messageBody := mail.BuildMessage()

	err := smtp.SendMail(os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"),
		smtp.PlainAuth("", mail.senderId, os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST")),
		mail.senderId, mail.toIds, []byte(messageBody))

	if err != nil {
		log.Printf("smtp error: %s", err)
	}

}
