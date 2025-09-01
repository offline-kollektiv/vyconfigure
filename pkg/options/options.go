package options

import "github.com/spf13/cobra"

type Options struct {
	Host            string
	ApiKey          string
	ConfigDirectory string
	Insecure        bool
	Timeout         int64
}

func GetOptions(cmd *cobra.Command, args []string) *Options {
	ConfigDirectory, _ := cmd.Flags().GetString("config-dir")
	Insecure, _ := cmd.Flags().GetBool("insecure")
	Timeout, _ := cmd.Flags().GetInt64("timeout")
	return &Options{
		Host:            "https://" + args[0],
		ApiKey:          args[1],
		ConfigDirectory: ConfigDirectory,
		Insecure:        Insecure,
		Timeout:         Timeout,
	}
}
