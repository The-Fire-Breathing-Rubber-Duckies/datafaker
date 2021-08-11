package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(entitiesCmd)
}

var entitiesCmd = &cobra.Command{
	Use:   "entities",
	Short: "Entities command",
	Long:  `Entities root command`,
}
