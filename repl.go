package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/masintxi/go_pokedex/internal/pokeapi"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command, ok := getcommands()[words[0]]
		if !ok {
			fmt.Println("Invalid command. Please try again.")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}

	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getcommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Page forward of the Pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Page back of the Pokemon locations",
			callback:    commandMapb,
		},
	}
}
