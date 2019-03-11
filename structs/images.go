package structs

// GameImages represents the different keys for images a Game can contain
type GameImages struct {
	Square    string `json:"square,omitempty"`
	Circle    string `json:"circle,omitempty"`
	Rectangle string `json:"rectangle,omitempty"`
}

// TournamentImages represents the different keys for images a Tournament can contain
type TournamentImages struct {
	Default   string `json:"default,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Banner    string `json:"banner,omitempty"`
	Square    string `json:"square,omitempty"`
	Fallback  bool   `json:"fallback"`
}

// TeamImages represents the different keys for images a Team can contain
type TeamImages struct {
	Default   string `json:"default,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Fallback  bool   `json:"fallback"`
}

// PlayerImages represents the different keys for images a Player can contain
type PlayerImages struct {
	Default   string `json:"default,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Fallback  bool   `json:"fallback"`
}

// RaceImages represents the different keys for images a Race can contain
type RaceImages struct {
	Default   string `json:"default,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

// StreamImages represents the different keys for images a Stream can contain
type StreamImages struct {
	UserLogo string `json:"user_logo,omitempty"`
	Preview  string `json:"preview,omitempty"`
}

// PlatformImages represents the different keys for images a Platform can contain
type PlatformImages struct {
	Default string `json:"default,omitempty"`
}

// CountryImages represents the different keys for images a Country can contain
type CountryImages struct {
	Default   string `json:"default,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}
