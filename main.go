package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/ryanguinchard/pokedexcli/internal/pokeapi"
)

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		// Get the command name from the first word of the input
		commandName := words[0]

		command, exists := cfg.commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v", err)
	}
}

func main() {
	nextLocationURL := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	cfg := &config{
		commands:        getCommands(),
		pokeapiClient:   pokeapi.NewClient(5 * time.Second),
		nextLocationURL: &nextLocationURL,
	}
	startREPL(cfg)

}
