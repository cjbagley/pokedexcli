package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cjbagley/pokedexcli/repl/cmds"
	"github.com/cjbagley/pokedexcli/utils"
)

const Prompt = "Pokédex > "
const HelpMsg = "Use 'help' for available command list.\n"

func StartRepl() {
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

		cmd.Callback()
		showPrompt()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func showPrompt() {
	fmt.Print(Prompt)
}
