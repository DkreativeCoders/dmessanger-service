package iservice

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/binding"
)

type IUserService interface {
	GetUser(id int) (*domain.User, error)
	GetAllUser() map[string]interface{}
	CreateUser(user domain.User) map[string]interface{}
	UpdatePassword(id int, request binding.UpdatePasswordRequest) binding.ResponseDto
}
