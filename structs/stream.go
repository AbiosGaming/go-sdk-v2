package structs

// StreamStruct represents a broadcast on a third party platform.
type StreamStruct struct {
	Id          int64              `json:"id,omitempty"`
	Username    string             `json:"username,omitempty"`
	DisplayName string             `json:"display_name,omitempty"`
	StatusText  string             `json:"status_text,omitempty"`
	ViewerCount int64              `json:"viewer_count"`
	Online      int64              `json:"online"`
	LastOnline  string             `json:"last_online,omitempty"`
	Images      StreamImagesStruct `json:"images,omitempty"`
	Url         string             `json:"url,omitempty"`
	Platform    PlatformStruct     `json:"platform,omitempty"`
}
