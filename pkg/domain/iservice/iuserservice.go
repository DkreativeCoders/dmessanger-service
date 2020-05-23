package iservice

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/binding"
)

type IUserService interface {
	GetUser(id int) (*domain.User, error)
	GetAllUser() map[string]interface{}
	CreateUser(user domain.User) (*domain.User, error)
	UpdatePassword(id int, request dto.UpdatePasswordRequest) error
}
