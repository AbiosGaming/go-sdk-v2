package structs

type DotaItemStruct struct {
	Image struct {
		Default   string `json:"default"`
		Thumbnail string `json:"thumbnail"`
	}
	Name string `json:"name"`
}
