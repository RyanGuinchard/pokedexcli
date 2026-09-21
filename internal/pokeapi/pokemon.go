package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(name string) (*Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	var pokemon Pokemon

	// check if the pokemon is already in cache
	if cachedPokemon, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(cachedPokemon, &pokemon); err != nil {
			return nil, err
		}
		return &pokemon, nil
	}

	// if not cached, fetch from the API
	response, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("pokemon not found: %s (status %d)", name, response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return nil, err
	}

	return &pokemon, nil
}
