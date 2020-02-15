package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configOutputCmd = &cobra.Command{
	Use:   "output",
	Short: "Output your configuration file.",
	Long: `Output your configuration file. 

The configuration file is stored at '~/.of/config.json'.`,

	Run: func(cmd *cobra.Command, args []string) {
		if err := config.Output(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	configCmd.AddCommand(configOutputCmd)
}
