package old

import (
	"errors"

	"gopkg.in/gomail.v2"
)

type Todo struct {
	Name       string
	Note       string
	Attachment string
}

func (todo *Todo) Send(email string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "omnifocus-cli@localhost.com")
	message.SetHeader("To", email)
	message.SetHeader("Subject", todo.Name)

	if todo.Note != "" {
		message.SetBody("text/plain", todo.Note)
	}

	if todo.Attachment != "" {
		message.Attach(todo.Attachment)
	}

	config := Configuration{}
	err := config.Parse()
	if err != nil {
		return err
	}

	if config.GmailEmail == "" {
		return errors.New("error: The gmail address in your configuration file is empty")
	}

	if config.GmailPassword == "" {
		return errors.New("error: The gmail password in your configuration file is empty")
	}

	dialer := gomail.Dialer{
		Host:     "smtp.gmail.com",
		Port:     465,
		SSL:      true,
		Username: config.GmailEmail,
		Password: config.GmailPassword,
	}

	err = dialer.DialAndSend(message)
	if err != nil {
		return err
	}

	return nil
}