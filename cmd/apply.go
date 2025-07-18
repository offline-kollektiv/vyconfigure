package cmd

import (
	"context"

	"github.com/offline-kollektiv/vyconfigure/pkg/api"
	"github.com/offline-kollektiv/vyconfigure/pkg/config"
	"github.com/offline-kollektiv/vyconfigure/pkg/convert"
	"github.com/offline-kollektiv/vyconfigure/pkg/options"
	r3diff "github.com/r3labs/diff/v3"
	"github.com/urfave/cli/v3"
)

func apply(ctx context.Context, cmd *cli.Command) error {
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
	changelog, err := r3diff.Diff(rc, lc)
	if err != nil {
		return err
	}

	var toDelete []string
	var toCreate []string
	if len(changelog) > 0 {
		for _, change := range changelog {
			if change.Type == "create" {
				toCreate = append(toCreate, change.To.(string))
			}
			if change.Type == "delete" {
				toDelete = append(toDelete, change.From.(string))
			}
		}
	} else {
		println("No changes to apply.")
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
