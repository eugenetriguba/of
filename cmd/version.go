package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the Omnifocus CLI",
	Long:  `See Omnifocus CLI's Version Number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Omnifocus CLI v0.2.0")
	},
}
