package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"github.com/grysha11/pokedex/api"
)

type CliCommand struct {
	Name		string
	Description	string
	Callback	func(cfg *api.Config, args []string) error
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
		"explore": {
			Name:			"explore",
			Description:	"Displays list of pokemons avaliable in given location area",
			Callback:		CommandExplore,
		},
		"catch": {
			Name:			"catch",
			Description:	"Gives a chance to catch a specific pokemon",
			Callback:		CommandCatch,
		},
	}
}

func CommandExit(cfg *api.Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *api.Config, args []string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, cmd := range Cmds {
		fmt.Printf(" %s: %s\n", cmd.Name, cmd.Description)
	}

	fmt.Printf("\n")

	return nil
}

func CommandMap(cfg *api.Config, args []string) error {
	locationData, err := api.GetLocationAreas(true, cfg)
	if err != nil {
		return err
	}

	for _, location := range locationData.Results {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}

func CommandMapB(cfg *api.Config, args []string) error {
	locationData, err := api.GetLocationAreas(false, cfg)
	if err != nil {
		return err
	}

	for _, location := range locationData.Results {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}

func CommandExplore(cfg *api.Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid argument: Try explore <location-area>")
	}

	pokemonData, err := api.GetLocationAreaPokemons(args[0], cfg)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", args[0])
	
	if len(pokemonData.PokemonEncounters) == 0 {
		fmt.Printf("None was found...\n")
	} else {
		fmt.Printf("Found pokemon:\n")
		for _, pokemonEncounter := range pokemonData.PokemonEncounters {
			fmt.Printf(" - %v\n", pokemonEncounter.Pokemon.Name)
		}
	}

	return nil
}

func CommandCatch(cfg *api.Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid argument: Try catch <pokemon-name>")
	}

	pokemon, err := api.GetPokemonData(args[0], cfg)
	if err != nil {
		return err
	}

	if _, ok := cfg.Pokedex[args[0]]; ok {
		fmt.Printf("You already have this pokemon!\n")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	const maxDifficulty = 500

	catchChance := maxDifficulty - pokemon.BaseExperience
	if catchChance < 0 {
		catchChance = 0
	}

	random := r.Intn(maxDifficulty)
	if random < catchChance {
		fmt.Printf("%v was caught!\n", args[0])
		cfg.Pokedex[args[0]] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", args[0])
	}

	return nil
} 
