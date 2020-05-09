package service

import (
	"github.com/danieloluwadare/dmessanger/oldstructure/interfaces/irepository"
	"github.com/danieloluwadare/dmessanger/oldstructure/interfaces/iservice"
	"github.com/danieloluwadare/dmessanger/oldstructure/models"
	"github.com/danieloluwadare/dmessanger/oldstructure/utils"
)

//INewService return an interface that's why Constrictor/Method name is preceded with I
func INewService(repository irepository.IUserRepository) iservice.IUserService {
	return service{repository}
}

type service struct {
	repository irepository.IUserRepository
}

//NewService Return the very instance of the class
func NewService(repository irepository.IUserRepository) *service {
	return &service{repository: repository}
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

//GetAllUser performs all necessary validation if need be and
//returns a map of data
func (s service) GetAllUser() map[string]interface{} {
	users := s.repository.FindAll()
	resp := utils.Message(true, "success")
	resp["data"] = users
	return resp
}

//
func (s service) GetUser(id int) (*models.User, error) {
	var user models.User
	return &user, nil

}
