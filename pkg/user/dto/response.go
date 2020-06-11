package dto

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/defaultresponse"
)

//swagger:model user-response-model
type UserResponseDto struct {
	// The ResponseData status
	//
	defaultresponse.ResponseData
	Data domain.User `json:"data"`
}
