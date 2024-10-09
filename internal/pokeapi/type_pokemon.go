package pokeapi

type Pokemon struct {
	LocationAreaEncounters string      `json:"location_area_encounters"`
	Cries                  Cries       `json:"cries"`
	Types                  []Type      `json:"types"`
	BaseExperience         int64       `json:"base_experience"`
	HeldItems              []HeldItem  `json:"held_items"`
	Weight                 int64       `json:"weight"`
	IsDefault              bool        `json:"is_default"`
	Sprites                Sprites     `json:"sprites"`
	PastTypes              []PastType  `json:"past_types"`
	Abilities              []Ability   `json:"abilities"`
	GameIndices            []GameIndex `json:"game_indices"`
	Species                Species     `json:"species"`
	Stats                  []Stat      `json:"stats"`
	Moves                  []Move      `json:"moves"`
	Name                   string      `json:"name"`
	ID                     int64       `json:"id"`
	Forms                  []Species   `json:"forms"`
	Height                 int64       `json:"height"`
	Order                  int64       `json:"order"`
}

type Ability struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int64   `json:"slot"`
	Ability  Species `json:"ability"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Cries struct {
	Legacy string `json:"legacy"`
	Latest string `json:"latest"`
}

type GameIndex struct {
	GameIndex int64   `json:"game_index"`
	Version   Species `json:"version"`
}

type HeldItem struct {
	Item           Species         `json:"item"`
	VersionDetails []VersionDetail `json:"version_details"`
}

type VersionDetail struct {
	Version Species `json:"version"`
	Rarity  int64   `json:"rarity"`
}

type Move struct {
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
	Move                Species              `json:"move"`
}

type VersionGroupDetail struct {
	LevelLearnedAt  int64   `json:"level_learned_at"`
	VersionGroup    Species `json:"version_group"`
	MoveLearnMethod Species `json:"move_learn_method"`
}

type PastType struct {
	Generation Species `json:"generation"`
	Types      []Type  `json:"types"`
}

type Type struct {
	Slot int64   `json:"slot"`
	Type Species `json:"type"`
}

type Sprites struct {
	Other        Other    `json:"other"`
	BackDefault  string   `json:"back_default"`
	FrontDefault string   `json:"front_default"`
	Versions     Versions `json:"versions"`
	BackShiny    string   `json:"back_shiny"`
	FrontShiny   string   `json:"front_shiny"`
}

type Other struct {
	DreamWorld      DreamWorld `json:"dream_world"`
	Showdown        Showdown   `json:"showdown"`
	OfficialArtwork Home       `json:"official-artwork"`
	Home            Home       `json:"home"`
}

type DreamWorld struct {
	FrontDefault string `json:"front_default"`
}

type Home struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type Showdown struct {
	BackDefault  string    `json:"back_default"`
	FrontDefault string    `json:"front_default"`
	BackShiny    string    `json:"back_shiny"`
	FrontShiny   string    `json:"front_shiny"`
	Animated     *Showdown `json:"animated,omitempty"`
}

type Versions struct {
	GenerationIii  GenerationIii   `json:"generation-iii"`
	GenerationIi   GenerationIi    `json:"generation-ii"`
	GenerationV    GenerationV     `json:"generation-v"`
	GenerationIv   GenerationIv    `json:"generation-iv"`
	GenerationVii  GenerationVii   `json:"generation-vii"`
	GenerationI    GenerationI     `json:"generation-i"`
	GenerationViii GenerationViii  `json:"generation-viii"`
	GenerationVi   map[string]Home `json:"generation-vi"`
}

type GenerationI struct {
	Yellow  RedBlue `json:"yellow"`
	RedBlue RedBlue `json:"red-blue"`
}

type RedBlue struct {
	FrontGray    string `json:"front_gray"`
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
}

type GenerationIi struct {
	Gold    Showdown `json:"gold"`
	Crystal Showdown `json:"crystal"`
	Silver  Showdown `json:"silver"`
}

type GenerationIii struct {
	FireredLeafgreen Showdown `json:"firered-leafgreen"`
	RubySapphire     Showdown `json:"ruby-sapphire"`
	Emerald          Home     `json:"emerald"`
}

type GenerationIv struct {
	Platinum            Showdown `json:"platinum"`
	DiamondPearl        Showdown `json:"diamond-pearl"`
	HeartgoldSoulsilver Showdown `json:"heartgold-soulsilver"`
}

type GenerationV struct {
	BlackWhite Showdown `json:"black-white"`
}

type GenerationVii struct {
	Icons             DreamWorld `json:"icons"`
	UltraSunUltraMoon Home       `json:"ultra-sun-ultra-moon"`
}

type GenerationViii struct {
	Icons DreamWorld `json:"icons"`
}

type Stat struct {
	Stat     Species `json:"stat"`
	BaseStat int64   `json:"base_stat"`
	Effort   int64   `json:"effort"`
}
