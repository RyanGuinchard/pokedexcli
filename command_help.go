package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
map: Displays a  list of map locations of the Pokemon world
mapb: Displays the previous list of map locations of the Pokemon world
explore <location>: Displays the Pokemon that can be found in a specific location.`)
	return nil
}
