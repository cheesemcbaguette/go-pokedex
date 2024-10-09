package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Location struct {
	GameIndex            int64                 `json:"game_index"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
	Name                 string                `json:"name"`
	Location             LocationClass         `json:"location"`
	ID                   int64                 `json:"id"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
}

type EncounterMethodRate struct {
	VersionDetails  []EncounterMethodRateVersionDetail `json:"version_details"`
	EncounterMethod LocationClass                      `json:"encounter_method"`
}

type LocationClass struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodRateVersionDetail struct {
	Rate    int64         `json:"rate"`
	Version LocationClass `json:"version"`
}

type Name struct {
	Name     string        `json:"name"`
	Language LocationClass `json:"language"`
}

type PokemonEncounter struct {
	Pokemon        LocationClass                   `json:"pokemon"`
	VersionDetails []PokemonEncounterVersionDetail `json:"version_details"`
}

type PokemonEncounterVersionDetail struct {
	MaxChance        int64             `json:"max_chance"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
	Version          LocationClass     `json:"version"`
}

type EncounterDetail struct {
	ConditionValues []interface{} `json:"condition_values"`
	Chance          int64         `json:"chance"`
	Method          LocationClass `json:"method"`
	MaxLevel        int64         `json:"max_level"`
	MinLevel        int64         `json:"min_level"`
}
