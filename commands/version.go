package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the Omnifocus Task Sender",
	Long:  `See Omnifocus Task Sender's Version Number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Omnifocus task sender v0.1.0")
	},
}
