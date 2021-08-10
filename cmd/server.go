package cmd

import (
	"github.com/spf13/cobra"
	"github.com/the-fire-breathing-duckies/datafaker/pkg"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Long:  "Run the DataFaker REST API server",
	Run: func(cmd *cobra.Command, args []string) {
		router := pkg.NewRouter()
		router.Logger.Fatal(router.Start(":1234"))
	},
}
