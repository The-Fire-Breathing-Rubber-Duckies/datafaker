package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile      string
	cfgFileName  = ".datafaker"
	cfgFileType  = "yaml"
	verboseLevel = 1
	silent       = false

	rootCmd = &cobra.Command{
		Use: "datafaker",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.datafaker.yaml)")
	rootCmd.PersistentFlags().IntVarP(&verboseLevel, "verbose-level", "v", verboseLevel, "verbose level")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "silent")
	rootCmd.PersistentFlags().String("host", "localhost", "host name for database connection")
	rootCmd.PersistentFlags().String("port", "5432", "port for database connection")
	rootCmd.PersistentFlags().String("user", "", "user for database connection")
	rootCmd.PersistentFlags().String("password", "", "password for database connection")
	rootCmd.PersistentFlags().String("dbname", "", "dbname for database connection")
	rootCmd.PersistentFlags().String("sslmode", "disable", "sslmode for database connection")
	viper.BindPFlag("hostname", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("dbname", rootCmd.PersistentFlags().Lookup("dbname"))
	viper.BindPFlag("sslmode", rootCmd.PersistentFlags().Lookup("sslmode"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType(cfgFileType)
		viper.SetConfigName(cfgFileName)
	}

	viper.AutomaticEnv()

	// Check silent
	if silent {
		verboseLevel = 0
	}

	if err := viper.ReadInConfig(); err == nil {
		if verboseLevel > 0 {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}
