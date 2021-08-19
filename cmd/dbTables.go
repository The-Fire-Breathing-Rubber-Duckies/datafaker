package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/datafakery/datafaker/pkg/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	dbCmd.AddCommand(dbTablesCmd)
}

var dbTablesCmd = &cobra.Command{
	Use:   "tables",
	Short: "Get tables",
	Long:  "Get table names from a db connection",
	Run: func(cmd *cobra.Command, args []string) {
		// Get connection config
		dbConnectParams := db.ConnectParams{
			Host:     viper.GetViper().GetString("hostname"),
			Port:     viper.GetViper().GetString("port"),
			User:     viper.GetViper().GetString("username"),
			Password: viper.GetViper().GetString("password"),
			Dbname:   viper.GetViper().GetString("dbname"),
			Sslmode:  viper.GetViper().GetString("sslmode"),
		}

		// Connect
		dbHandle := db.Connect(dbConnectParams)

		// Describe
		tables := db.GetTables(dbHandle)

		defer dbHandle.Close()

		out, err := json.Marshal(tables)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(out))
		}
	},
}
