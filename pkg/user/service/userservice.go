package service

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/dto"
	_ "github.com/DkreativeCoders/dmessanger-service/pkg/user/doc"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	"github.com/DkreativeCoders/dmessanger-service/pkg/utils"
	"errors"
)

//INewService return an interface that's why Constrictor/Method name is preceded with I
func INewService(repository irepository.IUserRepository) iservice.IUserService {
	return service{repository}
}

type service struct {
	repository irepository.IUserRepository

}

//perform validation on user and let UserRepository save user
func (s service) CreateUser(user domain.User) (*domain.User, error) {
	//user.Validate()
	if err := user.ValidateToError(); err!=nil {
		return nil, err
	}

	//s.repository.FindUserExist(user.Email)
	if found := s.repository.FindUserExist(user.Email); found{
		return nil,errors.New("user Already Exist with email")
	}


	newUser, err := s.repository.Save(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

//GetAllUser performs all necessary validation if need be and
//returns a map of data
func (s service) GetAllUser() map[string]interface{} {
	users := s.repository.FindAll()
	resp := utils.Message(true, "success")
	resp["data"] = users
	return resp
}

//
func (s service) GetUser(id int) (*domain.User, error) {
	var user domain.User
	return &user, nil

}

func (s service) UpdatePassword(id int, request dto.UpdatePasswordRequest) error {

	err := request.Validate()
	if err != nil {
		return err
	}

	user, err := s.repository.FindByID(id)



	if err != nil {
		return err
	}

	if user.Password == request.OldPassword {
		user.Password = request.NewPassword
		_, err := s.repository.Update(*user)
		if err != nil {
			return err
		}
		return nil
	}



	return errors.New("Incorrect password supplied.")

}
