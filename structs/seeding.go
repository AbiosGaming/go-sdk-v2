package structs

// SeedingStruct represents the seeding of two competitors in the related Series.
type SeedingStruct struct {
	Top    *int64 `json:"1,omitempty"` // Value is a roster-id
	Bottom *int64 `json:"2,omitempty"` // Value is a roster-id
}
