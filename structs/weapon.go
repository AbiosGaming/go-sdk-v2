package structs

// Weapon holds information about a CS:GO weapon.
type Weapon struct {
	Images struct {
		Small string `json:"small"`
	} `json:"images"`
	Name string `json:"name"`
}
