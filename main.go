package main

import (
	"time"

	"github.com/masintxi/go_pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokemonCaught: map[string]pokeapi.PokeInfo{},
	}
	startRepl(cfg)
}
