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

	"of/configuration"
	"of/fs"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "of",
	Short: "Of (Omnifocus) is a quick way to send tasks to your inbox",
	Long: `The Omnifocus task sender lets you quickly send tasks to your inbox.

The sent task will include the task name and optionally, a note or attachment.
Complete documentation is available at https://github.com/eugenetriguba/of`,
}

var config = configuration.Configuration{}

// Execute is the entry point to the commands that
// leverages cobra.Command.Execute() to parse the given arguments.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.Init)

	configPath, err := config.GetConfigFilePath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configExists, err := fs.FileExists(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if configExists {
		err := config.Parse()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}
}
