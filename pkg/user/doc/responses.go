package doc

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
)

// A TokenResponse .
//swagger:response tokenResponse
type UserResponseWrapper struct {
	// in: body
	Body domain.TokenResponse
}
