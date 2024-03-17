package cmds

import (
	"errors"
	"fmt"
)

func CatchCommand(config *Config, args ...string) error {
	if len(args) < 2 {
		return errors.New("please provide a Pokémon to try to catch")
	}

	if len(args) > 2 {
		return errors.New("please provide only one Pokémon")
	}

	pokémon, err := config.Client.GetPokemon(args[1])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokéball at %v\n", pokémon.Name)
	fmt.Printf("%v escaped!\n", pokémon.Name)

	return nil
}
