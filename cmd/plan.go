package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/offline-kollektiv/vyconfigure/pkg/api"
	"github.com/offline-kollektiv/vyconfigure/pkg/diff"
	"github.com/offline-kollektiv/vyconfigure/pkg/options"
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:     "plan <host> <api-key>",
	Aliases: []string{"p"},
	Short:   "Shows the diff between the current directory and VyOS instance",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		o := options.GetOptions(cmd, args)

		client, err := api.CreateClient(o)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		toDelete, toCreate, chg, err := diff.GetDiff(o, client)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		if chg {
			fmt.Println("Changes to be applied:")
			for _, change := range toDelete {
				color.Red("- delete " + change)
			}
			for _, change := range toCreate {
				color.Green("+ set " + change)
			}
		} else {
			fmt.Println("No changes to apply.")
		}
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
