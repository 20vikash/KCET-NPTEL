package gmail

import (
	helper "authentication/internal"
	"log"

	"gopkg.in/gomail.v2"
)

func SendMail() {
	m := gomail.NewMessage()

	m.SetHeader("From", "arisivikash@gmail.com")
	m.SetHeader("To", "arisivikash@gmail.com")
	m.SetHeader("Subject", "Hello")

	m.SetBody("text/html", "<html>Click <a href='http://localhost:8080/verify'>here</a> to activate your account</html>")

	d := gomail.NewDialer("smtp.gmail.com", 587, "arisivikash@gmail.com", helper.GetGmailAppPassword())

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Cannot send mail")
	}
}
