package cmds

type CliCommand struct {
	Name        string
	Description string
	Callback    func(args ...string) error
}

type Config struct {
	NextUrl     string
	PreviousUrl string
}

func GetCliCommands(config *Config) map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback: func(args ...string) error {
				return HelpCommand(config, args...)
			},
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 location areas. Each subsequent call will get the next 20 available",
			Callback: func(args ...string) error {
				return MapCommand(config, args...)
			},
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the last 20 location areas. Each subsequent call will get the last 20 available",
			Callback: func(args ...string) error {
				return MapbCommand(config, args...)
			},
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokédex",
			Callback: func(args ...string) error {
				return ExitCommand(config, args...)
			},
		},
	}
}
