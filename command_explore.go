package main

import (
	"fmt"

	"github.com/masintxi/go_pokedex/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a location name")
	}
	url := pokeapi.LocURL + args[0]

	data, err := cfg.pokeapiClient.GetPokemon(url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", data.Location.Name)
	fmt.Println("Found Pokemon:")
	for _, result := range data.PokemonEncounters {
		fmt.Println(" -", result.Pokemon.Name)
	}

	return nil

}
