package cmds

import (
	"github.com/cjbagley/pokedexcli/internal"
	"github.com/cjbagley/pokedexcli/utils"
)

func MapCommand(config *Config, args ...string) error {
	data, err := internal.GetLocations(config.NextUrl)
	if err != nil {
		return err
	}

	config.NextUrl, err = utils.GetEndpointFromUrl(data.Next)
	if err != nil {
		return err
	}

	config.PreviousUrl, err = utils.GetEndpointFromUrl(data.Previous)
	if err != nil {
		return err
	}
	data.PrintLocations()
	return nil
}
