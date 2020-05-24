package dto

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/defaultresponse"
)

//swagger:model CustomerResponseDtoModel
type UserResponseDto struct {
	// The ResponseData status
	//
	defaultresponse.ResponseData
	Data domain.Customer `json:"data"`
}