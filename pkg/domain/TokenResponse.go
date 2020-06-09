package domain

import "github.com/dgrijalva/jwt-go"

//User Entity  containing basic fields
//swagger:model token-response-model
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type LoginToken struct {
	Id          uint
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
	jwt.StandardClaims
}
