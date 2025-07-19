package options

import "github.com/urfave/cli/v3"

type Options struct {
	Host            string
	ApiKey          string
	ConfigDirectory string
	Insecure        bool
}

func GetOptions(cmd *cli.Command) *Options {
	return &Options{
		Host:            cmd.String("host"),
		ApiKey:          cmd.String("api-key"),
		ConfigDirectory: cmd.String("config-dir"),
		Insecure:        cmd.Bool("insecure"),
	}
}
