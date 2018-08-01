package structs

// SocialMediaAccountStruct represents a social media account for a Player or Team
type SocialMediaAccountStruct struct {
	Name string `json:"name,omitemtpy"`
	Slug string `json:"slug,omitempty"`
	Url  string `json:"url,omitmepty"`
}
