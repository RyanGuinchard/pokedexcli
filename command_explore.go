package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: explore <location>")
	}

	location := args[0]

	fmt.Printf("Exploring %s...\n", location)
	locationData, err := cfg.pokeapiClient.GetLocationAreaDetail(location)
	if err != nil {
		return fmt.Errorf("error fetching location data: %v", err)
	}

	fmt.Printf("Found Pokemon: \n")
	for _, encounter := range locationData.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
