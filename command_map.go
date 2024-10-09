package main

import (
	"fmt"
	"pokedex/internal/pokeapi"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	locationAreas, err := pokeapi.FetchLocationAreas(url)
	if err != nil {
		return err
	}

	// Print the location names
	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	// Update the Next and Previous URLs in the config
	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("No previous page available.")
		return nil
	}

	locationAreas, err := pokeapi.FetchLocationAreas(*cfg.Previous)
	if err != nil {
		return err
	}

	// Print the location names
	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	// Update the Next and Previous URLs in the config
	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous

	return nil
}
