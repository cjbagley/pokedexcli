package cmds

import (
	"github.com/cjbagley/pokedexcli/internal/types"
)

func PokedexCommand(config *Config, args ...string) error {
	pokedex := types.GetPokedex()
	pokedex.PrintCaughtPokemon()

	return nil
}
