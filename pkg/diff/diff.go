package diff

import (
	"github.com/offline-kollektiv/vyconfigure/pkg/api"
	"github.com/offline-kollektiv/vyconfigure/pkg/config"
	"github.com/offline-kollektiv/vyconfigure/pkg/convert"
	"github.com/offline-kollektiv/vyconfigure/pkg/options"
	r3diff "github.com/r3labs/diff/v3"
)

func GetDiff(o *options.Options, c *api.Client) (toDelete []string, toCreate []string, chg bool, err error) {

	// get remote config as cmds
	d, err := c.RetrieveJson()
	if err != nil {
		return
	}

	rc, _ := convert.JsonToCmds(d, "")

	// get local config as cmds
	lc, err := config.ReadAsCmds(o)
	if err != nil {
		return
	}

	// get diff
	changelog, err := r3diff.Diff(rc, lc)
	if err != nil {
		return
	}

	if len(changelog) > 0 {
		for _, change := range changelog {
			if change.Type == "create" {
				toCreate = append(toCreate, change.To.(string))
			}
			if change.Type == "delete" {
				toDelete = append(toDelete, change.From.(string))
			}
		}
		chg = true
	}

	return
}
