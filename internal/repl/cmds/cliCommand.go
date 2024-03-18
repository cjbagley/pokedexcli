package cmds

import "github.com/cjbagley/pokedexcli/internal/api"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config, args ...string) error
}

type Config struct {
	Client              api.Client
	LocationNextUrl     string
	LocationPreviousUrl string
}

func GetCliCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"catch": {
			Name:        "catch",
			Description: "Use this command with a given Pokémon to try and catch it! If you do, it will be added to your Pokédex",
			Callback:    CatchCommand,
		},
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
			Description: "Displays a help message.",
			Callback:    HelpCommand,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a previously caught Pokémon. The name of the Pokémon must be passed with this command.",
			Callback:    InspectCommand,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 location areas. Each subsequent call will get the next 20 available.",
			Callback:    MapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the last 20 location areas. Each subsequent call will get the last 20 available.",
			Callback:    MapbCommand,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Print details of all caught Pokémon in you Pokédex.",
			Callback:    PokedexCommand,
		},
	}
}
