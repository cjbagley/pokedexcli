package cmds

import (
	"fmt"

	"github.com/cjbagley/pokedexcli/repl/internal"
)

func MapCommand(config *Config, args ...string) error {
	data, err := internal.GetLocations(config.NextUrl)
	if err != nil {
		return err
	}

	config.NextUrl = data.Next
	config.PreviousUrl = data.Previous

	fmt.Println(data.Location[0])
	return nil
}
