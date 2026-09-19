package pokeapi

type LocationAreaResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []LocationArea
}

type LocationArea struct {
	Name string
	URL  string
}

type LocationAreaDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
