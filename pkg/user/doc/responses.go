package doc

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/dto"
)

//swagger:response userResponse
type UserResponseWrapper struct {
	// in: body
	Body struct {
		dto.UserResponseDto
	}
}