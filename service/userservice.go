package service

import (
	"github.com/danieloluwadare/dmessanger/interfaces/irepository"
	"github.com/danieloluwadare/dmessanger/interfaces/iservice"
	"github.com/danieloluwadare/dmessanger/models"
	"github.com/danieloluwadare/dmessanger/utils"
)

func NewService(repository irepository.IUserRepository) iservice.IUserService {
	return service{repository}
}

type service struct {
	repository irepository.IUserRepository
}

//perform validation on user and let UserRepository save user
func (s service) CreateUser(user models.User) map[string]interface{} {
	//user.Validate()
	if resp, ok := user.Validate(); !ok {
		return resp
	}
	 newUser, err := s.repository.Save(user)
	if err != nil{
		resp := utils.Message(false, "Error")
		resp["error_message"] = err
		return resp
	}
	resp := utils.Message(true, "success")
	resp["data"] = newUser
	return resp
}

func (s service) GetAllUser() map[string]interface{} {
	users := s.repository.FindAll()
	resp := utils.Message(true, "success")
	resp["data"] = users
	return resp
}

func (s service) GetUser(id int) (*models.User, error) {
	var user models.User
	return &user, nil

}
