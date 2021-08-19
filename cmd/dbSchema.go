package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/datafakery/datafaker/pkg/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	dbCmd.AddCommand(dbSchemaCmd)
}

var dbSchemaCmd = &cobra.Command{
	Use:   "schema [table]",
	Short: "Get table schema",
	Long:  `Get schema from table name`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tableName := args[0]

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
		tableSchema := db.DescribeTable(dbHandle, tableName)

		defer dbHandle.Close()

		out, err := json.Marshal(tableSchema)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(out))
		}
	},
}
