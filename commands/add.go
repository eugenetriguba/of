package commands

import (
	"fmt"
	"of/todo"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo into your omnifocus inbox",
	Long: `Add a todo into your omnifocus inbox
	
Usage:
	of add "my new todo" -n "cool note" -a "~/report.pdf"`,
	Run: func(cmd *cobra.Command, args []string) {
		newTodo.Name = args[0]
		err := newTodo.Send(config.MailDropEmail)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Successfully sent your todo!")
	},
}

var newTodo = todo.Todo{}

func init() {
	rootCmd.AddCommand(addCmd)
	configCmd.Flags().StringVarP(&newTodo.Note, "note", "n", "", "Additional note")
	configCmd.Flags().StringVarP(&newTodo.Attachment, "attachment", "a", "", "Absolute path to a file to attach to the todo")
}
