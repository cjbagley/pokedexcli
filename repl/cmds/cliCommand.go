package cmds

import "github.com/cjbagley/pokedexcli/internal"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config, args ...string) error
}

type Config struct {
	Client      internal.Client
	NextUrl     string
	PreviousUrl string
}

func GetCliCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokédex",
			Callback:    ExitCommand,
		},
		"explore": {
			Name:        "explore",
			Description: "Show a list of Pokémon within a given area. A location must be passed with this command.",
			Callback:    ExploreCommand,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    HelpCommand,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 location areas. Each subsequent call will get the next 20 available",
			Callback:    MapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the last 20 location areas. Each subsequent call will get the last 20 available",
			Callback:    MapbCommand,
		},
	}
}
