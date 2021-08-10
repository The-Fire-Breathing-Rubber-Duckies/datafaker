package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/the-fire-breathing-duckies/datafaker/pkg/db"
)

func init() {
	rootCmd.AddCommand(getSchemaCmd)
}

var getSchemaCmd = &cobra.Command{
	Use:   "schema [table]",
	Short: "Get table schema",
	Long:  `Get schema from table name`,
	Args:  cobra.ExactArgs(1),
	Run:   getSchema,
}

func getSchema(cmd *cobra.Command, args []string) {
	tableName := args[0]

	// Get connection config
	dbConnectParams := db.ConnectParams{
		Host:     viper.GetViper().GetString("hostname"),
		Port:     viper.GetViper().GetInt("port"),
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

	fmt.Printf("%v", tableSchema)
}
