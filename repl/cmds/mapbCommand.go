package cmds

import (
	"github.com/cjbagley/pokedexcli/utils"
)

func MapbCommand(config *Config, args ...string) error {
	data, err := config.Client.GetLocations(config.PreviousUrl)
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
