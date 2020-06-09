package doc

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/dto"
)

// swagger:parameters updatePassword
type UpdatePasswordRequestWrapper struct {
	// in: body
	Body dto.UpdatePasswordRequest

	// in: path
	UserID int
}

// swagger:parameters disableUser
type DisableUserRequestWrapper struct {
	// in: path
	UserID int
}

// swagger:parameters enableUser
type EnableUserRequestWrapper struct {
	// in: path
	UserID int
}

// swagger:parameters updatePassword
type LoginRequestWrapper struct {
	// in: body
	Body dto.LoginRequest
}