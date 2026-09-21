package main

import (
	"strings"

	"github.com/ryanguinchard/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

type config struct {
	commands        map[string]cliCommand
	pokeapiClient   *pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemon   map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
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
			description: "Displays a map of the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous map of the Pokemon world",
			callback:    commandMapB,
		},

		"explore": {
			name:        "explore",
			description: "Displays the Pokemon that can be found in a specific location. ",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokemon",
			callback:    commandCatch,
		},

		"inspect": {
			name:        "inspect",
			description: "Displays information about a caught Pokemon",
			callback:    commandInspect,
		},

		"pokedex": {
			name:        "pokedex",
			description: "Displays a list of caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	// Trim leading and trailing whitespace and lowercase the input
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	// Split the input into words based on whitespace
	words := strings.Fields(text)

	return words
}
