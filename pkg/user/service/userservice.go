package service

import (
	"errors"
	"fmt"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	_ "github.com/DkreativeCoders/dmessanger-service/pkg/user/doc"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/utils"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/mail"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/uuid"
)

//INewService return an interface that's why Constrictor/Method name is preceded with I
func INewService(repository irepository.IUserRepository, uuid uuid.IUuid, mailService mail.IMail) iservice.IUserService {
	return service{repository, uuid, mailService}
}

type service struct {
	repository irepository.IUserRepository
	uuid uuid.IUuid
	mailService mail.IMail
}

func (s service) EnableUser(id int) error {
	user, err := s.repository.FindByID(id)

	if err != nil {
		return err
	}

	if user.IsEnabled == false {
		user.IsEnabled = true
		_, err := s.repository.Update(*user)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (s service) DisableUser(id int) error {

	user, err := s.repository.FindByID(id)

	if err != nil {
		return err
	}

	if user.IsEnabled == true {
		fmt.Println(user)
		user.IsEnabled = false
		_, err := s.repository.Update(*user)
		if err != nil {
			return err
		}
		return nil
	} else {
		fmt.Println("user is enabled is false")
	}

	return nil
}

//perform validation on user and let UserRepository save user
func (s service) CreateUser(user domain.User) (*domain.User, error) {
	//user.Validate()
	if err := user.ValidateToError(); err != nil {
		return nil, err
	}

	//s.repository.FindUserExist(user.Email)
	if found := s.repository.FindUserExist(user.Email); found {
		return nil, errors.New("user Already Exist with email")
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

	if(user.Password == request.NewPassword) {
		return errors.New("Please select a new password")
	}

	if user.Password == request.OldPassword {
		user.Password = request.NewPassword
		_, err := s.repository.Update(*user)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("Incorrect password supplied")
}

func (s service) ForgotPassword(email string) error {

	var userExists = s.repository.FindUserExist(email)

	if !userExists {
		return errors.New("User not found")
	}

	var token = "fd"
	confirmationEmail := mail.NewEMailMessage(mail.ForgotPasswordSubject, ForgotPasswordMailBody(token), email, nil)


	_, err := s.mailService.SendEMail(*confirmationEmail)

	if err != nil {
		return errors.New("error occurred try to send mail, try again later")
	}


	return nil
}


func ForgotPasswordMailBody(link string) string {
	return "Please visit this link to reset your password. \n This links expires in an hour \n " + link + " \n Please ignore this mail if you didn't initiate this request."
}



