package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/grysha11/pokedex/internal/pokecache"
)

type Config struct {
	NextLocationArea	*string
	PrevLocationArea	*string
	PokeCache			*pokecache.Cache
	Pokedex				map[string]PokemonData
}

type LocationArea struct {
	Count    int		`json:"count"`
	Next     *string	`json:"next"`
	Previous *string	`json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaPokemons struct {
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonData struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func GetLocationAreas(forward bool, cfg *Config) (LocationArea, error) {
	var url string

	if !forward && cfg.PrevLocationArea == nil {
		fmt.Println("You are on the first page")
		return LocationArea{}, nil
	}

	if forward {
		url = *cfg.NextLocationArea
	} else {
		url = *cfg.PrevLocationArea
	}

	if data, ok := cfg.PokeCache.Get(url); ok {
		var location LocationArea
		err := json.Unmarshal(data, &location)
		if err != nil {
			return LocationArea{}, err
		}

		cfg.NextLocationArea = location.Next
		cfg.PrevLocationArea = location.Previous

		return location, nil
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("request was failed with status: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	cfg.PokeCache.Add(url, body)

	var location LocationArea
	err = json.Unmarshal(body, &location)
	if err != nil {
		return LocationArea{}, err
	}

	cfg.NextLocationArea = location.Next
	cfg.PrevLocationArea = location.Previous

	return location, nil
}

func GetLocationAreaPokemons(area string, cfg *Config) (LocationAreaPokemons, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + area + "/"

	if data, ok := cfg.PokeCache.Get(url); ok {
		var pokemons LocationAreaPokemons
		err := json.Unmarshal(data, &pokemons)
		if err != nil {
			return LocationAreaPokemons{}, err
		}

		return pokemons, nil
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaPokemons{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return LocationAreaPokemons{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return LocationAreaPokemons{}, fmt.Errorf("request was failed with status: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaPokemons{}, err
	}

	cfg.PokeCache.Add(url, body)

	var pokemons LocationAreaPokemons
	err = json.Unmarshal(body, &pokemons)
	if err != nil {
		return LocationAreaPokemons{}, err
	}

	return pokemons, nil
}

func GetPokemonData(name string, cfg *Config) (PokemonData, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name + "/"

	if data, ok := cfg.PokeCache.Get(url); ok {
		var pokemon PokemonData
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return PokemonData{}, err
		}

		return pokemon, nil
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonData{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return PokemonData{}, fmt.Errorf("request was failed with status: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonData{}, err
	}

	cfg.PokeCache.Add(url, body)

	var pokemon PokemonData
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return PokemonData{}, err
	}

	return pokemon, nil
}
