package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/the-fire-breathing-duckies/datafaker/pkg/entities"
)

var (
	domain string
	n      = 10
)

func init() {
	emailEntityCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "", "domain to use when generating email")
	emailEntityCmd.PersistentFlags().IntVarP(&n, "number", "n", n, "number of domains to generate")
	entitiesCmd.AddCommand(emailEntityCmd)
}

var emailEntityCmd = &cobra.Command{
	Use:   "emails",
	Short: "Email entity command",
	Long:  `Generate email entities`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		emails := entities.GetEmails(n, domain, "")

		out, err := json.Marshal(emails)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(out))
		}
	},
}
