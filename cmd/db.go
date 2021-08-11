package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dbCmd)
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "DB root command",
	Long:  "DB root command",
}
