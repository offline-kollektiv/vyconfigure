package cmd

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/offline-kollektiv/vyconfigure/pkg/api"
	"github.com/offline-kollektiv/vyconfigure/pkg/diff"
	"github.com/offline-kollektiv/vyconfigure/pkg/options"
	"github.com/urfave/cli/v3"
)

func plan(ctx context.Context, cmd *cli.Command) error {
	o := options.GetOptions(cmd)

	client, err := api.CreateClient(o)
	if err != nil {
		return err
	}

	toDelete, toCreate, chg, err := diff.GetDiff(o, client)
	if err != nil {
		return err
	}

	if chg {
		fmt.Println("Changes to be applied:")
		for _, change := range toDelete {
			color.Red("- delete " + change)
		}
		for _, change := range toCreate {
			color.Green("+ set " + change)
		}
	} else {
		fmt.Println("No changes to apply.")
	}

	return nil
}
