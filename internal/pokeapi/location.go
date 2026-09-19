package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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

func (c *Client) GetLocationAreaDetail(locationName string) (LocationAreaDetail, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", locationName)

	// check if the data is in the cache
	if cachedData, found := c.cache.Get(url); found {
		var detail LocationAreaDetail
		err := json.Unmarshal(cachedData, &detail)
		if err != nil {
			return LocationAreaDetail{}, err
		}
		return detail, nil
	}

	// If not in cache, fetch from the API
	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	defer res.Body.Close()

	// Add the response body to the cache
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	c.cache.Add(url, body)

	// Unmarshal the response body into the LocationAreaDetail struct
	var detail LocationAreaDetail
	err = json.Unmarshal(body, &detail)
	if err != nil {
		return LocationAreaDetail{}, err
	}
	return detail, nil
}
