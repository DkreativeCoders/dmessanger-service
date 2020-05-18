package binding

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
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


