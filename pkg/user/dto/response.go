package dto

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/defaultresponse"
)

//swagger:model UserResponseDtoModel
type UserResponseDto struct {
	defaultresponse.ResponseData
	Data domain.User `json:"data"`
}