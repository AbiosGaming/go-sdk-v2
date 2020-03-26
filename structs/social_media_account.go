package structs

// SocialMediaAccount represents a social media account for a Player or Team
type SocialMediaAccount struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
	Url  string `json:"url,omitempty"`
}
