package structs

// PlatformStruct represents a third party streaming platform.
type PlatformStruct struct {
	Id     int64                `json:"id,omitempty"`
	Name   string               `json:"name,omitempty"`
	Color  string               `json:"color,omitempty"`
	Images PlatformImagesStruct `json:"images,omitempty"`
}
