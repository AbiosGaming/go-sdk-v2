package structs

// WeaponStruct holds information about a CS:GO weapon.
type WeaponStruct struct {
	Images struct {
		Small string `json:"small"`
	} `json:"images"`
	Name string `json:"name"`
}
