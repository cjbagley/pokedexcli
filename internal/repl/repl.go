package repl

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/cjbagley/pokedexcli/internal/api"
	"github.com/cjbagley/pokedexcli/internal/repl/cmds"
	"github.com/cjbagley/pokedexcli/utils"
)

const Prompt = "Pokédex > "
const HelpMsg = "Use 'help' for available command list.\n"

func StartRepl() {
	config := cmds.Config{
		Client: api.NewClient(5*time.Second, 5*time.Minute),
	}
	cmds := cmds.GetCliCommands()
	scanner := bufio.NewScanner(os.Stdin)

	showPrompt()
	for scanner.Scan() {
		input := utils.SanitisePromptInput(scanner.Text())
		if len(input) == 0 {
			fmt.Printf("No command provided. %v", HelpMsg)
			showPrompt()
			continue
		}

		cmd, ok := cmds[input[0]]
		if !ok {
			fmt.Printf("Command not recognised. %v", HelpMsg)
			showPrompt()
			continue
		}

		err := cmd.Callback(&config, input...)
		if err != nil {
			DisplayError(err)
		}
		showPrompt()
	}
	if err := scanner.Err(); err != nil {
		ExitWithError(err)
	}
}

func showPrompt() {
	fmt.Print(Prompt)
}

func DisplayError(err error) {
	fmt.Printf("Error: %v\n", err.Error())
}

func ExitWithError(err error) {
	DisplayError(err)
	os.Exit(1)
}
