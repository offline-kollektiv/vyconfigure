package cmd

import (
	"context"

	"github.com/offline-kollektiv/vyconfigure/pkg/api"
	"github.com/offline-kollektiv/vyconfigure/pkg/config"
	"github.com/offline-kollektiv/vyconfigure/pkg/options"
	"github.com/urfave/cli/v3"
)

func sync(ctx context.Context, cmd *cli.Command) error {
	o := options.GetOptions(cmd)

	client, err := api.CreateClient(o)
	if err != nil {
		return err
	}

	d, err := client.Retrieve()
	if err != nil {
		return err
	}

	err = config.Write(d, o)
	if err != nil {
		return err
	}

	return nil
}
