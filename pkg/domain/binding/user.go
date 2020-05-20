package binding

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"errors"
)


//swagger:model
type UserResponse struct {
	ResponseDto
	Data domain.User `json:"data"`
}

//swagger:model
type UpdatePasswordRequest struct { 
	OldPassword string `json:"OldPassword"`
	NewPassword string `json:"NewPassword"`
	ConfirmNewPassword string `json:"ConfirmNewPassword"`
}

//Validate All the fields for User
func (request *UpdatePasswordRequest) Validate() error {

	if request.NewPassword != request.ConfirmNewPassword {
		return errors.New("Passwords don't match")
	}

	return nil
}


