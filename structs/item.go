package structs

type DotaItem struct {
	Image struct {
		Default   string `json:"default"`
		Thumbnail string `json:"thumbnail"`
	} `json:"image"`
	Name string `json:"name"`
}

type LolItem struct {
	Name       string `json:"name"`
	ExternalId int64  `json:"external_id"`
	Image      struct {
		Default   string `json:"default"`
		Thumbnail string `json:"thumbnail"`
	} `json:"image"`
}
