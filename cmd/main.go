package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().String("config-dir", ".", "Directory where config is stored.")
	rootCmd.PersistentFlags().Bool("insecure", false, "Whether to skip verifying the SSL certificate.")
	rootCmd.Version = GetVersion()
	rootCmd.SetVersionTemplate("{{.DisplayName}} {{.Version}}\n")
}

var rootCmd = &cobra.Command{
	Use:   "vyconfigure",
	Short: "vyconfigure - Declarative configuration for VyOS.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
