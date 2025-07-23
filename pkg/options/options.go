package options

import "github.com/spf13/cobra"

type Options struct {
	Host            string
	ApiKey          string
	ConfigDirectory string
	Insecure        bool
}

func GetOptions(cmd *cobra.Command, args []string) *Options {
	return &Options{
		Host:            "https://" + args[0],
		ApiKey:          args[1],
		ConfigDirectory: cmd.Flag("config-dir").Value.String(),
		Insecure:        cmd.Flag("insecure").Value.String() == "true",
	}
}
