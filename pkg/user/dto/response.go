package dto

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/utils"
)

//swagger:model UserResponseDtoModel
type UserResponseDto struct {
	utils.ResponseDto
	Data domain.User `json:"data"`
}