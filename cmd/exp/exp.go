package main

import "github.com/tombracho/ch-lenslocked/models"

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 2525
	username = "2bffcb431ac475"
	passwrd  = "01ab3e159d8d53"
)

func main() {
	email := models.Email{
		From:      "from@example.com",
		To:        "artyomsonyx@gmail.com",
		Subject:   "Test mail",
		Plaintext: "Body of the mail",
		HTML:      `<h1>Hello there buddy!</h1><p>This is the email</p><p>Hope you enjoy it</p>`,
	}

	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: passwrd,
	})

	es.Send(email)

}
