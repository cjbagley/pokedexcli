package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cjbagley/pokedexcli/utils"
)

const Prompt = "Pokédex > "

func StartRepl() {
	cmds := GetCliCommands()
	scanner := bufio.NewScanner(os.Stdin)

	showPrompt()
	for scanner.Scan() {
		input := utils.SanitisePromptInput(scanner.Text())

		cmd, ok := cmds[input[0]]
		if !ok {
			fmt.Println("Command not recognised. Use 'help' for available command list")
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
