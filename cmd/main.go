package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Version, commit and date are stamped at link time via -ldflags -X.
// They have placeholder defaults so `go build` from a clean
// checkout still produces a runnable binary without the Makefile.
var (
	Version = "0.0.0-dev"
	Commit  = "unknown"
	Date    = "unknown"
)

func init() {
	rootCmd.PersistentFlags().String("config-dir", ".", "Directory where config is stored.")
	rootCmd.PersistentFlags().Bool("insecure", false, "Whether to skip verifying the SSL certificate.")
	rootCmd.PersistentFlags().Int64("timeout", 30, "Timeout in seconds for API requests.")
}

var rootCmd = &cobra.Command{
	Use:   "vyconfigure",
	Short: "vyconfigure - Declarative configuration for VyOS.",
	Version: fmt.Sprintf("%s (commit %s, built %s)", Version, Commit, Date),
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
