package domain

//{"access_token":"Z_1QUVC5M_EOCESISKW8AQ","expires_in":7200,"scope":"read","token_type":"Bearer"}
type TokenResponse struct {
	AccessToken       string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	Scope string `json:"scope"`
	TokenType string  `json:"token_type"`
}