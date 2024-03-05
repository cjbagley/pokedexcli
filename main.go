package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cjbagley/pokedexcli/cmds"
	"github.com/cjbagley/pokedexcli/utils"
)

const Prompt = "Pokedex > "

func main() {
	cmds := cmds.GetCliCommands()
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
