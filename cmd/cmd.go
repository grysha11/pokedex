package cmd

import (
	"fmt"
	"os"
	"github.com/grysha11/pokedex/api"
)

type CliCommand struct {
	Name		string
	Description	string
	Callback	func(cfg *api.Config) error
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
		"map": {
			Name:			"map",
			Description:	"Displays the names of next 20 location areas",
			Callback:		CommandMap,
		},
		"mapb": {
			Name:			"mapb",
			Description:	"Displays the names of previous 20 location areas",
			Callback:		CommandMapB,
		},
	}
}

func CommandExit(cfg *api.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *api.Config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, cmd := range Cmds {
		fmt.Printf(" %s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func CommandMap(cfg *api.Config) error {
	locationData, err := api.GetLocationAreas(true, cfg)
	if err != nil {
		return err
	}

	for _, location := range locationData.Results {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}

func CommandMapB(cfg *api.Config) error {
	locationData, err := api.GetLocationAreas(false, cfg)
	if err != nil {
		return err
	}

	for _, location := range locationData.Results {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}
