// Package todo contains a todo type that represents a todo sent to Omnifocus.
package todo

import (
	"github.com/jordan-wright/email"
	errorFmt "github.com/pkg/errors"
	"net/smtp"
	"of/configuration"
)

// Todo represents a todo that will be sent into Omnifocus.
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
func (todo *Todo) Send() error {
	config := configuration.Configuration{}
	err := config.Parse()
	if err != nil {
		return errorFmt.Wrap(err, "Parsing the configuration file failed")
	}

	message, err := todo.constructEmail(config.MailDropEmail)
	if err != nil {
		return err
	}

	err = message.Send(
		"smtp.gmail.com:587",
		smtp.PlainAuth(
			"",
			config.GmailUsername,
			config.GmailPassword,
			"smtp.gmail.com",
		),
	)
	if err != nil {
		return errorFmt.Wrap(err, "Sending the todo failed")
	}

	return nil
}

func (todo *Todo) constructEmail(emailAddress string) (*email.Email, error) {
	message := email.NewEmail()
	message.From = "omnifocus-cli@localhost.com"
	message.To = []string{emailAddress}
	message.Subject = todo.Name

	if todo.Note != "" {
		message.Text = []byte(todo.Note)
	}

	if todo.Attachment != "" {
		_, err := message.AttachFile(todo.Attachment)
		if err != nil {
			return nil, errorFmt.Wrap(err, "error: could not attach the specified file to the email")
		}
	}

	return message, nil
}
