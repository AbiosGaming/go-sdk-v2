package structs

// AccessTokenStruct holds information about an access token.
type AccessTokenStruct struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}
