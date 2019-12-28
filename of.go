// Omnifocus (OF) CLI
//
// The Omnifocus CLI allows you to quickly send items to your omnifocus inbox.
//
// You'll first want to configure your mail drop email using the -email flag. After that,
// you can get started. Anything after 'of' will be parsed as the todo item name. You can
// then also add a additional note or attachment using a flag.
//
// Usage:
//   $ of -email fake.mail.drop.email@sync.omnigroup.com
//   > Successfully set your omnifocus mail drop email to fake.mail.drop.email@sync.omnigroup.com
//   $ of "Take dog for a walk" -note "Make sure he is on his leash" -attachment ~/dogs/walk-schedule.pdf
//   > Successfully sent your todo!

package main

import (
	"flag"
	"fmt"
	"of/configuration"
	"of/todo"
	"os"
)

func main() {
	email := flag.String("email", "", "Configure the Omnifocus mail drop email you'd like to use.")
	note := flag.String("note", "", "Set a note for the todo item.")
	attachment := flag.String("attachment", "", "Add an attachment along with the todo item. This should be an absolute path to the item.")
	flag.Parse()

	config := configuration.Configuration{}

	if *email != "" {
		config.MailDropEmail = *email
		err := config.Save()
		if err != nil {
			fmt.Println("An error occurred while trying to set your mail drop email: ", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully set your omnifocus mail drop email to %s\n", *email)
	}

	err := config.Parse()
	if err != nil {
		fmt.Println("An error occurred while trying to parse the configuration file: ", err)
		os.Exit(1)
	}

	if config.MailDropEmail == "" {
		fmt.Println("You haven't set an email in your configuration file! You'll want to do that first.")
		os.Exit(1)
	}

	newTodo := todo.Todo{Name: flag.Arg(0), Note: *note, Attachment: *attachment}
	fmt.Println(flag.NFlag())
	fmt.Println(newTodo)
	fmt.Println(*note)
	fmt.Println(*attachment)
	err = newTodo.Send(config.MailDropEmail)
	if err != nil {
		fmt.Println("An error occurred while trying to send the email: ", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully sent your todo!")
}
