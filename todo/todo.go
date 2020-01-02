// Package todo contains a todo type that represents a todo
// sent to Omnifocus.
package todo

import (
	"errors"

	"of/configuration"

	errorFmt "github.com/pkg/errors"
	"gopkg.in/gomail.v2"
)

// Todo represents a todo that will be
// sent into Omnifocus.
type Todo struct {
	Name       string
	Note       string
	Attachment string
}

// Send uses the fields of your todo to send off an
// email to the specified email address. The email is
// sent using the gmail username and password from your
// configuration file.
//
// The email is sent in the following way:
//   - The name of the todo is used as the subject line
//   - The note of the todo is used as the email body
//   - The attachment of the todo is, well, used as an attachment on the email.
//   - The email will be from "omnifocus-cli@localhost.com".
func (todo *Todo) Send(emailAddress string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "omnifocus-cli@localhost.com")
	message.SetHeader("To", emailAddress)
	message.SetHeader("Subject", todo.Name)

	if todo.Note != "" {
		message.SetBody("text/plain", todo.Note)
	}

	if todo.Attachment != "" {
		message.Attach(todo.Attachment)
	}

	config := configuration.Configuration{}
	err := config.Parse()
	if err != nil {
		return err
	}

	if config.GmailUsername == "" {
		return errors.New("error: The gmail username in your configuration file is empty")
	}

	if config.GmailPassword == "" {
		return errors.New("error: The gmail password in your configuration file is empty")
	}

	dialer := gomail.Dialer{
		Host:     "smtp.gmail.com",
		Port:     465,
		SSL:      true,
		Username: config.GmailUsername,
		Password: config.GmailPassword,
	}

	err = dialer.DialAndSend(message)
	if err != nil {
		return errorFmt.Wrap(err, "Sending the todo failed")
	}

	return nil
}
