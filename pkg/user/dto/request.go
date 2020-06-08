package dto

import (
	"errors"
)

//swagger:model update-password-request-model
type UpdatePasswordRequest struct {
	OldPassword        string `json:"OldPassword"`
	NewPassword        string `json:"NewPassword"`
	ConfirmNewPassword string `json:"ConfirmNewPassword"`
}

//Validate All the fields for User
func (request *UpdatePasswordRequest) Validate() error {

	if request.NewPassword != request.ConfirmNewPassword {
		return errors.New("Passwords don't match")
	}

	return nil
}

//{"access_token":"Z_1QUVC5M_EOCESISKW8AQ","expires_in":7200,"scope":"read","token_type":"Bearer"}
type LoginRequest struct {
	Email       string `json:"email"`
	Password    string `json:"-"`
}

func (request *LoginRequest) Validate() error {

	//if request.NewPassword != request.ConfirmNewPassword {
	//	return errors.New("Passwords don't match")
	//}
	//
	return nil
}