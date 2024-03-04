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

	fmt.Print(Prompt)
	for scanner.Scan() {
		input := r.ReplaceAllString(scanner.Text(), "")

		fmt.Println(input)

		cmd, ok := cmds[input]
		if ok {
			cmd.Callback()
		}
		fmt.Print(Prompt)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
