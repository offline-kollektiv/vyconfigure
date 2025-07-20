package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func version(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("vyconfigure", cmd.Version)
	os.Exit(0)
	return nil
}
