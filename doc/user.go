package doc

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/binding"
)

//swagger:response userResponse
type UserResponseWrapper struct {
	Body struct {
		binding.UserResponse
	}
}

// swagger:parameters updatePassword
type UpdatePasswordRequestWrapper struct {
	// in: body
	Body struct {
		binding.UpdatePasswordRequest
	}
}

// swagger:parameters updatePassword
type UserID struct {
	// in: path
	UserID int
}
