package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/the-fire-breathing-duckies/datafaker/pkg/db"
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

	dbParams := db.ConnectParams{
		"localhost",
		5321,
		"admin",
		"admin",
		"test",
	}
	db.connect(dbParams)
}
