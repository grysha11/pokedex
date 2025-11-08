package repl

import (
	"fmt"
	"bufio"
	"os"
	"github.com/grysha11/pokedex/util"
	"github.com/grysha11/pokedex/cmd"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := util.CleanInput(input)

		if len(words) < 1 {
			continue
		}

		command, ok := cmd.Cmds[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.Callback()
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
		}
	}
}