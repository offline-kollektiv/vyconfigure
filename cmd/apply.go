package cmd

import (
	"context"
	"fmt"

	"github.com/offline-kollektiv/vyconfigure/pkg/api"
	"github.com/offline-kollektiv/vyconfigure/pkg/convert"
	"github.com/offline-kollektiv/vyconfigure/pkg/diff"
	"github.com/offline-kollektiv/vyconfigure/pkg/options"
	"github.com/urfave/cli/v3"
)

func apply(ctx context.Context, cmd *cli.Command) error {
	o := options.GetOptions(cmd)

	client, err := api.CreateClient(o)
	if err != nil {
		return err
	}

	toDelete, toCreate, chg, err := diff.GetDiff(o, client)
	if err != nil {
		return err
	}

	if !chg {
		fmt.Println("No changes to apply.")
		return nil
	}

	dc := convert.CmdsToData(toDelete, "delete")
	cc := convert.CmdsToData(toCreate, "set")

	cmds := append(dc, cc...)
	err = client.Configure(cmds)
	if err != nil {
		return err
	}

	return nil
}
