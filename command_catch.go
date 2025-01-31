package main

import (
	"fmt"
	"math/rand"

	"github.com/masintxi/go_pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a pokemon name")
	}
	pokemon := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	url := pokeapi.PokeURL + pokemon

	data, err := cfg.pokeapiClient.CatchPokemon(url)
	if err != nil {
		if err.Error() == "404" {
			fmt.Printf("There was no \"%s\" here! Try with another pokemon\n", pokemon)
		}
		return err
	}

	baseExp := data.BaseExperience
	randNum := rand.Intn(baseExp)
	if baseExp-randNum > baseExp/3 {
		fmt.Printf("%s escaped! (%v vs %v)\n", pokemon, randNum, baseExp)
	} else {
		fmt.Printf("You caught a %s! (%v vs %v)\n", pokemon, randNum, baseExp)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokemonCaught[pokemon] = data
	}

	return nil
}
