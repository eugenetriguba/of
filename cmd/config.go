package cmd

import (
	"fmt"
	"os"

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
			return
		}

		if err := config.Save(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Configuration updated!")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(
		&config.MailDropEmail,
		"maildrop", "m", "",
		"Set your Omnifocus Mail Drop Email",
	)
	configCmd.Flags().StringVarP(
		&config.GmailUsername,
		"username", "u", "",
		"Set your Gmail Username",
	)
	configCmd.Flags().StringVarP(
		&config.GmailPassword,
		"password", "p", "",
		"Set your Gmail Password",
	)
	configCmd.Flags().StringVarP(
		&config.ApiKey,
		"apikey", "a", "",
		"Set your gmail api key.",
	)
}
