package dto

import (
	"errors"
)

//swagger:model updatePasswordRequestModel
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


