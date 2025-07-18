package cmd

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func version(ctx context.Context, cmd *cli.Command) error {
	println(appVersion)
	os.Exit(0)
	return nil
}
