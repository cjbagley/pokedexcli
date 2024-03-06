package cmds

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
	Config      *Config
}

type Config struct {
	NextUrl     string
	PreviousUrl string
}

func GetCliCommands() map[string]CliCommand {
	return map[string]CliCommand{
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
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokédex",
			Callback:    ExitCommand,
		},
	}
}
