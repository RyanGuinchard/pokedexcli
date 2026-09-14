package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// loop for user input until the program is exited
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

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v", err)
	}
}
