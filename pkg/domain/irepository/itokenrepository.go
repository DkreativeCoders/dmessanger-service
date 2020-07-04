package irepository

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
)

// IUserRepository
// Define your basic crud operation in this interface

type ITokenRepository interface {
	Create(token domain.Token) (*domain.Token, error)
	FindByUserId(userId int) (*domain.Customer, error)
	FindByToken(tk string) (*domain.Token, error)
	FindByOtp(otp string) (*domain.Token, error)
	UpdateToken(tk domain.Token) (*domain.Token, error)
}
