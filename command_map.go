package main

import (
	"fmt"

	"github.com/masintxi/go_pokedex/internal/pokeapi"
)

const locURL = "https://pokeapi.co/api/v2/location-area/"

func commandMap(cfg *config) error {
	url := locURL
	if cfg.Next != nil {
		url = *cfg.Next
	}

	if url == "" {
		fmt.Println("No more results.")
		return nil
	}

	printGetResults(cfg, url)
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("You are at the first page.")
		return nil
	}
	url := *cfg.Previous

	printGetResults(cfg, url)
	return nil
}

func printGetResults(cfg *config, url string) error {
	var data pokeapi.PokeMap

	data, err := cfg.pokeapiClient.GetLocations(url)
	if err != nil {
		return err
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	if len(data.Results) == 0 {
		fmt.Println("No location areas to display")
	} else {
		for _, result := range data.Results {
			fmt.Println(result.Name)
		}
	}
	fmt.Println()
	return nil
}
