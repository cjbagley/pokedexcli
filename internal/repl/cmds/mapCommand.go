package cmds

import (
	"github.com/cjbagley/pokedexcli/utils"
)

func MapCommand(config *Config, args ...string) error {
	list, err := config.Client.GetLocationList(config.LocationNextUrl)
	if err != nil {
		return err
	}

	config.LocationNextUrl, err = utils.GetEndpointFromUrl(list.Next)
	if err != nil {
		return err
	}

	config.LocationPreviousUrl, err = utils.GetEndpointFromUrl(list.Previous)
	if err != nil {
		return err
	}

	list.Print()
	return nil
}
