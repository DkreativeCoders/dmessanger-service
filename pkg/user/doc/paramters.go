package doc

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/dto"
)

// swagger:parameters updatePassword
type UpdatePasswordRequestWrapper struct {
	// in: body
	Body struct {
		dto.UpdatePasswordRequest
	}
}

// swagger:parameters updatePassword
type UserID struct {
	// in: path
	UserID int
}