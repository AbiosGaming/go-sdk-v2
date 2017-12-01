package structs

// GameImagesStruct represents the different keys for images a GameStruct can contain
type GameImagesStruct struct {
	Square    string `json:"square,omitempty"`
	Circle    string `json:"circle,omitempty"`
	Rectangle string `json:"rectangle,omitempty"`
}

// TournamentImagesStruct represents the different keys for images a TournamentStruct can contain
type TournamentImagesStruct struct {
	Default   string `json:"default,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Banner    string `json:"banner,omitempty"`
	Square    string `json:"square,omitempty"`
	Fallback  bool   `json:"fallback"`
}

// TeamImagesStruct represents the different keys for images a TeamStruct can contain
type TeamImagesStruct struct {
	Default   string `json:"default,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Fallback  bool   `json:"fallback"`
}

// PlayerImagesStruct represents the different keys for images a PlayerStruct can contain
type PlayerImagesStruct struct {
	Default   string `json:"default,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Fallback  bool   `json:"fallback"`
}

// RaceImagesStruct represents the different keys for images a RaceStruct can contain
type RaceImagesStruct struct {
	Default   string `json:"default,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

// StreamImagesStruct represents the different keys for images a StreamStruct can contain
type StreamImagesStruct struct {
	UserLogo string `json:"user_logo,omitempty"`
	Preview  string `json:"preview,omitempty"`
}

// PlatformImagesStruct represents the different keys for images a PlatformStruct can contain
type PlatformImagesStruct struct {
	Default string `json:"default,omitempty"`
}

// CountryImagesStruct represents the different keys for images a CountryStruct can contain
type CountryImagesStruct struct {
	Default   string `json:"default,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}
