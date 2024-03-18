package cmds

import (
	"errors"
	"fmt"

	"github.com/cjbagley/pokedexcli/internal/types"
)

func InspectCommand(config *Config, args ...string) error {
	if len(args) < 2 {
		return errors.New("please provide a Pokémon to inspect")
	}

	if len(args) > 2 {
		return errors.New("please provide only one Pokémon")
	}

	name := args[1]
	pokedex := types.GetPokedex()
	pokémon, ok := pokedex.GetByName(name)

	if !ok {
		return errors.New("pokémon has not yet been caught!")
	}

	fmt.Println(pokémon.Url)

	return nil
}
