package cmds

import (
	"github.com/cjbagley/pokedexcli/repl/internal"
)

func MapCommand(config *Config, args ...string) error {
	if config.NextUrl == "" {
		config.NextUrl = "location"
	}

	_, err := internal.GetLocations(config.NextUrl)
	if err != nil {
		return err
	}
	return nil
}
