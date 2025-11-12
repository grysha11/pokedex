package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	NextLocationArea	*string
	PrevLocationArea	*string
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

	dec := json.NewDecoder(resp.Body)
	var location LocationArea
	err = dec.Decode(&location)
	if err != nil {
		return LocationArea{}, err
	}

	cfg.NextLocationArea = location.Next
	cfg.PrevLocationArea = location.Previous

	return location, nil
}