package structs

type RoleStruct struct {
	Name string  `json:"name"`
	From string  `json:"from"`
	To   *string `json:"to"`
}
