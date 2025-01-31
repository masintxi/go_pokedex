package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a pokemon name")
	}
	pokemon := args[0]

	data, ok := cfg.pokemonCaught[pokemon]
	if !ok {
		fmt.Printf("You haven't caught %s yet!\n", pokemon)
		return nil
	}

	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %v\n", data.Height)
	fmt.Printf("Weight: %v\n", data.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range data.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typ := range data.Types {
		fmt.Printf("  -%s\n", typ.Type.Name)
	}
	return nil
}
