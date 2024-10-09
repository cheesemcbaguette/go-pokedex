package pokeapi

import (
	"encoding/json"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
}

type LocationAreaResponse struct {
	Results  []LocationArea `json:"results"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
}

// FetchLocationAreas fetches location areas from the API
func FetchLocationAreas(url string) (*LocationAreaResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locationAreas LocationAreaResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationAreas); err != nil {
		return nil, err
	}

	return &locationAreas, nil
}
