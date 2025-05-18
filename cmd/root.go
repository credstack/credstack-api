/*
Copyright © 2025 Steven A. Zaluk
*/

package cmd

import (
	"fmt"
	"github.com/stevezaluk/credstack-api/api"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "credstack-api",
	Short: "",
	Long:  `RESTful API for CredStack IDP`,
	Run: func(cmd *cobra.Command, args []string) {
		api := api.FromConfig()

		err := api.Start(viper.GetInt("port"))
		if err != nil {
			fmt.Println("\nAPI has exited due to an error: ", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.credstack/config.json)")

	rootCmd.Flags().IntP("port", "p", 8080, "The default port that the API is going to listen for requests at")
	viper.BindPFlags(rootCmd.Flags())
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + "/.credstack")
		viper.SetConfigType("json")
		viper.SetConfigName("config.json")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to read config file: ", err.Error())
		os.Exit(1)
	}
}
