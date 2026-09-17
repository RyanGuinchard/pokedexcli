package main

import "fmt"

func commandMapB(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return fmt.Errorf("you're on the first page")
	}

	locationData, err := cfg.pokeapiClient.GetLocationAreas(*cfg.prevLocationURL)
	if err != nil {
		return err
	}

	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

	cfg.prevLocationURL = locationData.Previous
	cfg.nextLocationURL = locationData.Next

	return nil
}
