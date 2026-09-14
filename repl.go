package main

import "strings"

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
