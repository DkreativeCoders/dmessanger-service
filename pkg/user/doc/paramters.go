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
