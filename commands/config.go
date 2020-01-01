package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Modify your configuration file.",
	Long: `Modify your configuration file. 

The configuration file is stored at '~/.of/config.json'. 
You can update your omnifocus mail drop email, gmail username, 
and gmail password.`,

	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 {
			fmt.Println(cmd.Help())
		}

		config.Save()
		fmt.Println("Configuration updated!")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&config.MailDropEmail, "maildrop", "m", "", "Omnifocus Mail Drop Email")
	configCmd.Flags().StringVarP(&config.GmailUsername, "username", "u", "", "Gmail Username")
	configCmd.Flags().StringVarP(&config.GmailPassword, "password", "p", "", "Gmail Password")
}
