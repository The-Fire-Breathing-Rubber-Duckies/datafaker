package cmd

import (
	"fmt"
	"strconv"

	"github.com/datafakery/datafaker/pkg"
	"github.com/spf13/cobra"
)

var port int

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVar(&port, "server-port", 1234, "server port")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run REST server",
	Long:  "Run the DataFaker REST API server",
	Run: func(cmd *cobra.Command, args []string) {
		host := ":" + strconv.Itoa(port)
		fmt.Printf("Starting server on %s", host)
		router := pkg.NewRouter()
		router.Logger.Fatal(router.Start(host))
	},
}
