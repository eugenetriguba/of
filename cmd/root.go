package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userLicense string
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "of",
	Short: "Of (Omnifocus) is a quick way to send tasks to your inbox",
	Long: `'of' lets you quickly send tasks with their name and an optional 
			note or attachment to your inbox. Complete documentation is available 
			at https://github.com/eugenetriguba/of`,
}

func Execute() {
	err := rootCmd.Execute()
	abortIf(err)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.of.json")
	rootCmd.PersistentFlags().StringP("author", "a", "Eugene Triguba", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "MIT", "Name of license for the project")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Eugene Triguba <eugenetriguba@gmail.com>")
	viper.SetDefault("license", "MIT")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		abortIf(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".of")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	abortIf(err)
}

func abortIf(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
