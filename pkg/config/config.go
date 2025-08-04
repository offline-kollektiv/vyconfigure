package config

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/offline-kollektiv/vyconfigure/pkg/convert"
	"github.com/offline-kollektiv/vyconfigure/pkg/options"
	"sigs.k8s.io/yaml"
)

// Write writes existing vyos config to the local filesystem
func Write(data map[string]interface{}, o *options.Options) error {
	for k := range data {
		y, err := yaml.Marshal(data[k])
		if err != nil {
			return err
		}

		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		p := path.Join(wd, o.ConfigDirectory, k+".yaml")
		err = os.WriteFile(p, y, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReadAsCmds reads all yaml configuration and converts it into vyos "set" commands
func ReadAsCmds(o *options.Options) ([]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	p := path.Join(wd, o.ConfigDirectory)
	files, err := os.ReadDir(p)
	if err != nil {
		return nil, err
	}
	if !hasYamlFile(files) {
		log.Fatal("No configuration files found.")
	}

	var res []string
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".yaml") {
			fp := path.Join(wd, o.ConfigDirectory, f.Name())

			c, err := os.ReadFile(fp)
			if err != nil {
				return nil, err
			}

			configPath := strings.TrimSuffix(f.Name(), ".yaml")
			cmds, err := convert.YamlToCmds(c, configPath+" ")
			if err != nil {
				return nil, err
			}
			res = append(res, cmds...)
		}
	}

	return res, nil
}

// Returns true if at least one entry ends with ".yaml"
func hasYamlFile(files []os.DirEntry) bool {
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".yaml") {
			return true
		}
	}
	return false
}
