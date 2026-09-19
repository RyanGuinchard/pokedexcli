package main

import "fmt"

func commandMap(cfg *config, args ...string) error {
	locationData, err := cfg.pokeapiClient.GetLocationAreas(*cfg.nextLocationURL)
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
