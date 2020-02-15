package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/eugenetriguba/of/todo"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo into your omnifocus inbox",
	Long: `Add a todo into your omnifocus inbox

You can add the task name and optionally, a note or attachment.
`,
	Example: `of add "my new todo" -n "cool note" -a "~/report.pdf"`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newTodo.Name = args[0]
		err := newTodo.Send()
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
	addCmd.Flags().StringVarP(
		&newTodo.Note,
		"note", "n", "",
		"Additional information to go into the note section of the todo",
	)
	addCmd.Flags().StringVarP(
		&newTodo.Attachment,
		"attachment", "a", "",
		"Absolute path to a file to attach to the todo",
	)
}
