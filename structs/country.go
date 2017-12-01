package structs

/* CountryStruct hold information about a country, nationality or language
 * associated with a resource.
 */
type CountryStruct struct {
	Name      string              `json:"name,omitempty"`
	ShortName string              `json:"short_name,omitempty"`
	Images    CountryImagesStruct `json:"images,omitempty"`
}
