package cmds

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func GetCliCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "displays a help message",
			Callback:    helpCommand,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    exitCommand,
		},
	}
}

func helpCommand() error {
	fmt.Println("Help message")
	return nil
}

func exitCommand() error {
	defer os.Exit(0)
	return nil
}
