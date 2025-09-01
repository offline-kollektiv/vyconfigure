package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/offline-kollektiv/vyconfigure/pkg/api"
	"github.com/offline-kollektiv/vyconfigure/pkg/convert"
	"github.com/offline-kollektiv/vyconfigure/pkg/diff"
	"github.com/offline-kollektiv/vyconfigure/pkg/options"
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:     "apply <host> <api-key>",
	Aliases: []string{"a"},
	Short:   "Applies the current configuration.",
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

		if !chg {
			fmt.Println("No changes to apply.")
			log.Fatal(err)
			os.Exit(1)
		}

		dc := convert.CmdsToData(toDelete, "delete")
		cc := convert.CmdsToData(toCreate, "set")

		cmds := append(dc, cc...)
		save, _ := cmd.Flags().GetBool("save")
		err = client.Configure(cmds, save)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	},
}

func init() {
	applyCmd.Flags().Bool("save", true, "Save the configuration after applying changes.")
	rootCmd.AddCommand(applyCmd)
}
