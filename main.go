package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/cjbagley/pokedexcli/cmds"
)

const Prompt = "Pokedex > "

func main() {
	cmds := cmds.GetCliCommands()
	scanner := bufio.NewScanner(os.Stdin)
	r, _ := regexp.Compile("[^A-Za-z]")

	showPrompt()
	for scanner.Scan() {
		input := r.ReplaceAllString(scanner.Text(), "")

		cmd, ok := cmds[input]
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
