// Package commands implements all the commands in this CLI.
//
// You can get the current version using `of version`, add
// a todo with `of add`, and customize configuration with `of config`.
package commands

import (
	"fmt"
	"os"

	"of/configuration"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "of",
	Short: "Of (Omnifocus) is a quick way to send tasks to your inbox",
	Long: "The Omnifocus task sender lets you quickly send tasks to your inbox. \n" +
		"The sent task will include the task name and optionally, a note or attachment. \n\n" +
		"Complete documentation is available at https://github.com/eugenetriguba/of",
}

var config = configuration.Configuration{}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.Init)

	if err := config.Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
