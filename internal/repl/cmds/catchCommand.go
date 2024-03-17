package cmds

import (
	"errors"
	"fmt"

	"github.com/cjbagley/pokedexcli/internal/types"
	"github.com/cjbagley/pokedexcli/utils"
)

func CatchCommand(config *Config, args ...string) error {
	if len(args) < 2 {
		return errors.New("please provide a Pokémon to try to catch")
	}

	if len(args) > 2 {
		return errors.New("please provide only one Pokémon")
	}

	pokedex := types.GetPokedex()

	pokémon, err := config.Client.GetPokemon(args[1])
	if err != nil {
		return err
	}

	if pokedex.IsCaught(pokémon) {
		return errors.New("pokémon has already been caught")
	}

	fmt.Printf("Throwing a Pokéball at %v\n", pokémon.Name)

	ball := utils.RandomThresholdCheck{
		Value: int(pokémon.BaseExperience),
	}
	isCaught := utils.PassesRandomThreshholdCheck(ball)

	if !isCaught {
		fmt.Printf("%v escaped!\n", pokémon.Name)
		return nil
	}

	fmt.Printf("%v caught!\n", pokémon.Name)
	pokedex.Add(pokémon)
	fmt.Printf("%v added to Pokédex!\n", pokémon.Name)
	return nil
}
