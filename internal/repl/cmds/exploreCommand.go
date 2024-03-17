package cmds

import (
	"errors"
)

func ExploreCommand(config *Config, args ...string) error {
	if len(args) < 2 {
		return errors.New("please provide an area")
	}

	if len(args) > 2 {
		return errors.New("please provide only one area")
	}

	location, err := config.Client.GetLocation(args[1])
	if err != nil {
		return err
	}

	location.PrintAreaPokemon()

	return nil
}
