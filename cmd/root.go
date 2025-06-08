/*
Copyright © 2025 Steven A. Zaluk
*/

package cmd

import (
	"fmt"
	"github.com/stevezaluk/credstack-api/api"
	"github.com/stevezaluk/credstack-api/api/handlers/management"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "credstack-api",
	Short: "",
	Long:  `RESTful API for CredStack IDP`,
	Run: func(cmd *cobra.Command, args []string) {
		api.App = api.New()

		/*
			Management Routes
		*/
		api.App.Get("/management/application", management.GetApplicationHandler)

		err := api.Start(viper.GetInt("port"))
		if err != nil {
			fmt.Println("Failed to start API:", err)
			os.Exit(1)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		err := api.Stop() // this won't work
		if err != nil {
			fmt.Println("Failed to stop API:", err)
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

	/*
		Database - Provides options that control how CredStack connects to MongoDB
	*/
	rootCmd.Flags().String("mongo.hostname", "127.0.0.1", "The hostname of your running MongoDB server")
	rootCmd.Flags().Int("mongo.port", 27017, "The port of your running MongoDB server")
	rootCmd.Flags().Int("mongo.connection_timeout", 15, "The number of seconds that MongoDB should wait before closing the connection")
	rootCmd.Flags().Bool("mongo.use_authentication", true, "If set to true, then authentication options will be evaluated")
	rootCmd.Flags().String("mongo.default_database", "credstack", "The default database that credstack will initialize in")
	rootCmd.Flags().String("mongo.authentication_database", "admin", "The default database in MongoDB that provides authentication")
	rootCmd.Flags().String("mongo.username", "", "The username that credstack will use for authentication with MongoDB")
	rootCmd.Flags().String("mongo.password", "", "The password that credstack will use for authentication with MongoDB")

	/*
		Log - Provides options that control how logging is handled
	*/
	rootCmd.Flags().String("log.level", "", "The level of logging to use. Can be one of: debug, warn, info. Defaults to info")
	rootCmd.Flags().String("log.path", "/var/log/credstack", "The directory to write log files too")
	rootCmd.Flags().Bool("log.use_file_logging", false, "If set to true, then log files will be written. Otherwise, only STDOUT logging will be used")

	/*
		Credential - Provides options that control how user credentials are hashed
	*/
	rootCmd.Flags().Uint32("argon.time", 1, "The number of iterations that will be made when hashing passwords with Argon2id")
	rootCmd.Flags().Uint32("argon.memory", 16*1024, "The amount of memory that argon can consume while hashing passwords")
	rootCmd.Flags().Uint8("argon.threads", 1, "The number of goroutines that argon can use while hashing passwords")
	rootCmd.Flags().Uint32("argon.key_length", 16, "The length that passwords will be hashed to")
	rootCmd.Flags().Uint32("argon.salt_length", 32, "The length that a salt will be generated to")
	rootCmd.Flags().Uint32("argon.min_secret_length", 12, "The minimum length requirement of plaintext user credentials")
	rootCmd.Flags().Uint32("argon.max_secret_length", 48, "The maximum length requirement of plaintext user credentials")

	err := viper.BindPFlags(rootCmd.Flags())
	if err != nil {
		fmt.Println("Failed to bind command flags: ", err.Error())
		os.Exit(1)
	}
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

	viper.SetEnvPrefix("CREDSTACK")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No config file was detected. Either default values or environmental variables will be used")
	}
}
