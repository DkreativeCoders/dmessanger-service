package mocks

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/stretchr/testify/mock"
)

// UserRepository is a mock implementation of the IUserRepository interface
type UserRepository struct {
	mock.Mock
}

func (ur *UserRepository) FindByID(id int) (*domain.User, error) {
	args := ur.Called(id)

	var r0 *domain.User
	if userFunc, ok := args.Get(0).(func(int) *domain.User); ok {
		r0 = userFunc(id)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*domain.User)
		}
	}

	var r1 error
	if errFunc, ok := args.Get(1).(func(int) error); ok {
		r1 = errFunc(id)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (ur *UserRepository) FindAll() []domain.User {
	args := ur.Called()

	var r0 []domain.User
	if userListFunc, ok := args.Get(0).(func() []domain.User); ok {
		r0 = userListFunc()
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]domain.User)
		}
	}

	return r0
}

func (ur *UserRepository) Save(user domain.User) (*domain.User, error) {
	args := ur.Called(user)

	var r0 *domain.User
	if userFunc, ok := args.Get(0).(func(domain.User) *domain.User); ok {
		r0 = userFunc(user)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*domain.User)
		}
	}

	var r1 error
	if errFunc, ok := args.Get(1).(func(domain.User) error); ok {
		r1 = errFunc(user)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (ur *UserRepository) Update(user domain.User) (*domain.User, error) {
	args := ur.Called(user)

	var r0 *domain.User
	if userFunc, ok := args.Get(0).(func(domain.User) *domain.User); ok {
		r0 = userFunc(user)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).(*domain.User)
		}
	}

	var r1 error
	if errFunc, ok := args.Get(1).(func(domain.User) error); ok {
		r1 = errFunc(user)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (ur *UserRepository) FindByEmail(email string) (*domain.User, error) {
	return nil, nil
}
func (ur *UserRepository) FindUserExist(email string) bool {
	return true
}
