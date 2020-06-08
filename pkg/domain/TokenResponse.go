package domain

import "github.com/dgrijalva/jwt-go"

//{"access_token":"Z_1QUVC5M_EOCESISKW8AQ","expires_in":7200,"scope":"read","token_type":"Bearer"}
type TokenResponse struct {
	AccessToken       string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	Scope string `json:"scope"`
	TokenType string  `json:"token_type"`
}

type LoginToken struct {
	Id uint
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
	jwt.StandardClaims
}