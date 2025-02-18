package gmail

import (
	helper "authentication/internal"
	"log"
	"net/smtp"
)

func SendMail() {
	auth := smtp.PlainAuth("", "arisivikash@gmail.com", helper.GetGmailAppPassword(), "smtp.gmail.com")

	to := []string{"arisivikash@gmail.com"}

	msg := []byte("To: arisivikash@gmail.com\r\n" +
		"Subject: Hello\r\n" +
		"\r\n" +
		"Hi\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "arisivikash@gmail.com", to, msg)

	if err != nil {
		log.Fatal("Error sending email")
	}
}
