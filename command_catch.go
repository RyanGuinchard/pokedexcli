package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	threshold := 50

	baseExp := pokemon.BaseExperience
	if baseExp <= 0 {
		// Fallback default experience if the API returns 0 or unmarshals empty
		baseExp = 50
	}

	randNum := rand.Intn(baseExp)

	if randNum > threshold {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	//Store in Pokedex on success
	cfg.caughtPokemon[pokemon.Name] = *pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}
