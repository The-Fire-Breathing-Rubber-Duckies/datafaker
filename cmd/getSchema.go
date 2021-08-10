package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/the-fire-breathing-duckies/datafaker/pkg"
)

func init() {
	rootCmd.AddCommand(getSchemaCmd)
}

var getSchemaCmd = &cobra.Command{
	Use:   "schema [table]",
	Short: "Get table schema",
	Long:  `Get schema from table name`,
	Run:   getSchema,
}

func getSchema(cmd *cobra.Command, args []string) {
	fmt.Print("Hello")

	dbParams := pkg.ConnectParams{
		Host:     "localhost",
		Port:     5321,
		User:     "admin",
		Password: "admin",
		Dbname:   "test",
	}
	pkg.Connect(dbParams)
}
