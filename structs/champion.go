package structs

type Champion struct {
	Name       string `json:"name"`
	ExternalId int64  `json:"external_id"`
	Images     struct {
		Default   string `json:"default"`
		Thumbnail string `json:"thumbnail"`
	} `json:"images"`
}
