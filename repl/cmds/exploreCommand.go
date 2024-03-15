package cmds

import (
	"github.com/cjbagley/pokedexcli/utils"
)

func ExploreCommand(config *Config, args ...string) error {
	data, err := config.Client.GetLocationList(config.LocationNextUrl)
	if err != nil {
		return err
	}

	config.LocationNextUrl, err = utils.GetEndpointFromUrl(data.Next)
	if err != nil {
		return err
	}

	config.LocationPreviousUrl, err = utils.GetEndpointFromUrl(data.Previous)
	if err != nil {
		return err
	}
	return nil
}
