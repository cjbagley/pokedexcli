package cmds

import (
	"errors"

	"github.com/cjbagley/pokedexcli/utils"
)

func ExploreCommand(config *Config, args ...string) error {
	if len(args) < 2 {
		return errors.New("please provide an area")
	}

	if len(args) > 2 {
		return errors.New("please provide only one area")
	}

	area := utils.SanitisePromptInput(args[1])
	location, err := config.Client.GetLocation(area[0])
	if err != nil {
		return err
	}

	location.PrintAreaPokemon()

	return nil
}
