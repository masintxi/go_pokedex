package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokemonCaught) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokemonCaught {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
