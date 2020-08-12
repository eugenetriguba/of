// Package commands implements all the commands in this CLI.
//
// Commands
//
// 'of add': Add a todo to your omnifocus inbox.
// 'of config': Modify the configuration file using flags.
// 'of config output': Output the current configuration file to stdout.
// 'of version': Print out the current version of the cli.
package commands

import (
	"fmt"
	"os"

	"github.com/eugenetriguba/of/configuration"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "of",
	Short: "Quickly send tasks to your Omnifocus inbox",
	Long: `Quickly send tasks to your Omnifocus inbox.

The sent task will include the task name and optionally, a note or attachment.
Complete documentation is available at https://github.com/eugenetriguba/of`,
}

var config = configuration.Configuration{}

// Execute is the entry point to the commands. It
// leverages cobra.Command.Execute() to parse the given arguments.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	if err := config.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := config.Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
