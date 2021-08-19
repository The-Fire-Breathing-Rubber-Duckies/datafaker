package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/datafakery/datafaker/pkg/entities"
	"github.com/spf13/cobra"
)

func init() {
	entitiesCmd.AddCommand(namesEntityCmd)
}

var namesEntityCmd = &cobra.Command{
	Use:   "names [num=10]",
	Short: "Names entity command",
	Long:  `Names entity command`,
	Args:  cobra.MaximumNArgs(1),
	Run:   namesEntity,
}

func namesEntity(cmd *cobra.Command, args []string) {
	var err error
	n := 10

	if len(args) > 0 {
		n, err = strconv.Atoi(args[0])

		if err != nil {
			panic("Invalid number of names")
		}
	}

	names := entities.GetNames(n)

	out, err := json.Marshal(names)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}
}
