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

// swagger:parameters userAuthentication
type LoginRequestWrapper struct {
	// in: body
	Body dto.LoginRequest
}

// swagger:parameters forgotPassword
type UserEmailParam struct {
	// in: path
	Email string
}

// swagger:parameters resetPassword
type ResetPasswordRequestWrapper struct {
	// in: body
	Body dto.ResetPasswordRequest

	// in: path
	Token string
}
