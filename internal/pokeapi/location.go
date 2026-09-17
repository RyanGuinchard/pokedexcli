package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetLocationAreas(url string) (LocationAreaResponse, error) {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	var data LocationAreaResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return data, nil
}
