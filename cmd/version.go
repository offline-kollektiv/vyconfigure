package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vyconfigure", GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func GetVersion() (ret string) {
	if b, ok := debug.ReadBuildInfo(); ok && len(b.Main.Version) > 0 {
		ret = b.Main.Version
	} else {
		ret = "unknown"
	}
	return
}
