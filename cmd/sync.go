package cmd

import (
	"log"
	"os"

	"github.com/offline-kollektiv/vyconfigure/pkg/api"
	"github.com/offline-kollektiv/vyconfigure/pkg/config"
	"github.com/offline-kollektiv/vyconfigure/pkg/options"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:     "sync <host> <api-key>",
	Aliases: []string{"s"},
	Short:   "Syncs configuration to the current directory through the HTTP API.",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		o := options.GetOptions(cmd, args)

		client, err := api.CreateClient(o)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		d, err := client.Retrieve()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		err = config.Write(d, o)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
