package cmd

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/offline-kollektiv/vyconfigure/pkg/api"
	"github.com/offline-kollektiv/vyconfigure/pkg/config"
	"github.com/offline-kollektiv/vyconfigure/pkg/convert"
	"github.com/offline-kollektiv/vyconfigure/pkg/options"
	diff "github.com/r3labs/diff/v3"
	"github.com/urfave/cli/v3"
)

func plan(ctx context.Context, cmd *cli.Command) error {
	o := options.GetOptions(cmd)

	// get remote config as cmds
	client, err := api.CreateClient(o)
	if err != nil {
		return err
	}

	d, err := client.RetrieveJson()
	if err != nil {
		return err
	}

	rc, _ := convert.JsonToCmds(d, "")

	// get local config as cmds
	lc, err := config.ReadAsCmds(o)
	if err != nil {
		return err
	}

	// get diff
	changelog, err := diff.Diff(rc, lc)
	if err != nil {
		return err
	}

	if len(changelog) > 0 {
		fmt.Println("Changes to be applied:")
		for _, change := range changelog {
			if change.Type == "create" {
				color.Green("+ set " + change.To.(string))
			}
			if change.Type == "delete" {
				color.Red("- delete " + change.From.(string))
			}
		}
	} else {
		fmt.Println("No changes to apply.")
	}

	return nil
}
