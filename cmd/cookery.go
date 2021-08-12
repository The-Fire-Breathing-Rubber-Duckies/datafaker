package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/the-fire-breathing-duckies/datafaker/pkg"
	"github.com/the-fire-breathing-duckies/datafaker/pkg/db"
)

func init() {
	rootCmd.AddCommand(cookeryCmd)
}

type TableSchema struct {
	Table  string
	Schema db.TableSchema
}

var cookeryCmd = &cobra.Command{
	Use:   "cookery [file]",
	Short: "Populate database",
	Long:  "Read cookery and populate database",
	Args:  cobra.ExactArgs(1),
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

		var schemas []TableSchema

		for _, table := range tables {
			schema := TableSchema{
				Table:  table,
				Schema: db.DescribeTable(dbHandle, table),
			}

			schemas = append(schemas, schema)
		}

		pkg.ParseCookery(dbHandle, args[0])

		defer dbHandle.Close()

		// out, err := json.Marshal(schemas)
		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println(string(out))
		// }
	},
}
