package cmd

import (
	"context"
	"os"
	"runtime/debug"

	"net/mail"

	"github.com/offline-kollektiv/vyconfigure/pkg/options"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"
)

func GetVersion() (ret string) {
	if b, ok := debug.ReadBuildInfo(); ok && len(b.Main.Version) > 0 {
		ret = b.Main.Version
	} else {
		ret = "unknown"
	}
	return
}

func Run() {
	o := options.Options{}

	app := &cli.Command{
		Name:                  "vyconfigure",
		Usage:                 "Declarative configuration for VyOS.",
		Version:               GetVersion(),
		EnableShellCompletion: true,
		Action:                version,

		Authors: []any{
			mail.Address{Name: "Timo Schirmer", Address: "timoschirmer@users.noreply.github.com"},
			mail.Address{Name: "Peter Lehmann", Address: "peterablehmann@users.noreply.github.com"},
		},

		Flags: []cli.Flag{
			&cli.StringFlag{Destination: &o.Host, Name: "host", Usage: "The hostname of the VyOS HTTP API.", Sources: cli.EnvVars("VYCONFIGURE_HOST")},
			&cli.StringFlag{Destination: &o.ApiKey, Name: "api-key", Usage: "API key for the HTTP API.", Sources: cli.EnvVars("VYCONFIGURE_API_KEY")},
			&cli.StringFlag{Destination: &o.ConfigDirectory, Name: "config-dir", Value: ".", Usage: "Directory where config is stored.", Sources: cli.EnvVars("VYCONFIGURE_CONFIG_DIR")},
			&cli.BoolFlag{Destination: &o.Insecure, Name: "insecure", Usage: "Whether to skip verifying the SSL certificate.", Sources: cli.EnvVars("VYCONFIGURE_INSECURE")},
		},

		Commands: []*cli.Command{
			{
				Name: "version", Aliases: []string{"v"}, Usage: "prints the current version.",
				Action: version,
			},
			{
				Name: "apply", Aliases: []string{"a"}, Usage: "applies the current configuration.",
				Action: apply,
			},
			{
				Name: "sync", Aliases: []string{"s"}, Usage: "syncs configuration to the current directory through the HTTP API.",
				Action: sync,
			},
			{
				Name: "plan", Aliases: []string{"d"}, Usage: "shows the diff between the current directory and VyOS instance",
				Action: plan,
			},
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
