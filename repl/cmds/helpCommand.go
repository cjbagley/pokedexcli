package cmds

import "fmt"

func HelpCommand(config *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokédex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCliCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()

	return nil
}
