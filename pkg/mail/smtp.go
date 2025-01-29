package mail

import (
	"os"

	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
)

var _ = godotenv.Load()

var host = os.Getenv("EMAIL_HOST")

var defaultSubject = os.Getenv("EMAIL_DEFAULT_SUBJECT")
var sender = os.Getenv("EMAIL_SENDER")

func Send(recipient string, subject string, body string) {
	// Create a new message
	message := gomail.NewMessage()

	// Set email headers with multiple recipients
	message.SetHeader("From", sender)
	message.SetHeader("To", recipient)
	message.SetHeader("Subject", defaultSubject)

	// Set email body
	message.SetBody("text/html", body)

	// Set up the SMTP mail server dialer
	dialer := gomail.NewDialer(
		host,
		587,
		os.Getenv("MAILTRAP_SANDBOX_USER"),
		os.Getenv("MAILTRAP_SANDBOX_PASSWORD"),
	)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		panic(err)
	}
}
