package cmd

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Name		string
	Description	string
	Callback	func() error
}

var Cmds map[string]CliCommand

func init() {
	Cmds = map[string]CliCommand {
		"exit": {
			Name:			"exit",
			Description:	"Exit the Pokedex",
			Callback:		CommandExit,
		},
		"help": {
			Name:			"help",
			Description:	"Displays a help message",
			Callback:		CommandHelp,
		},
	}
}

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, cmd := range Cmds {
		fmt.Printf(" %s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}