package service

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/binding"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
)

//INewService return an interface that's why Constrictor/Method name is preceded with I
func INewService(repository irepository.ICustomerRepository,userService iservice.IUserService ) iservice.ICustomerService{
	return service{repository, userService}
}

type service struct {
	repository irepository.ICustomerRepository
	userService iservice.IUserService

}

func (s service) CreateUser(request dto.CustomerRequest) (*domain.Customer, error){
	 err := request.Validate()
	if err != nil {
		return nil, err
		//return *defaultresponse.NewResponseDto(false, err.Error(),nil)
	}

	//FirstName   string `json:"firstName"`
	//LastName    string `json:"lastName"`
	//Age         string `json:"age"`
	//Email       string `json:"email"`
	//PhoneNumber string `json:"phoneNumber"`
	//Password    string `json:"-"`
	//Address     string `json:"address"`
	user := domain.User{
		FirstName: request.FirstName,
		LastName: request.LastName,
		Age: request.Age,
		Email: request.Age,
		PhoneNumber: request.PhoneNumber,
		Password: request.Password,
		Address: request.Address,
	}

	if err := user.ValidateToError(); err!=nil {
		return nil, err
	}

	newUser, err := s.userService.CreateUser(user)

	if err != nil {
		return nil, err
	}

	customer := domain.Customer{
		UserId: newUser.ID,
	}

	s.repository.Save(customer)

	panic("implement me")
}

//perform validation on user and let UserRepository save user
func (s service) UpdatePassword(id int, request binding.UpdatePasswordRequest) binding.ResponseDto {
	// swagger:operation PUT /api/vi/users/update-password/{UserID} updatePassword
	//
	// Updates a user's password
	// ---
	// responses:
	//   default:
	//     "$ref": "#/responses/responseDto"

	err := request.Validate()
	if err != nil {
		return *binding.NewResponseDto(false, err.Error())
	}

	user, err := s.repository.FindByID(id)

	if err != nil {
		return *binding.NewResponseDto(false, err.Error())
	}

	if user.Password == request.OldPassword {
		user.Password = request.NewPassword
		_, err := s.repository.Update(*user)
		if err != nil {
			return *binding.NewResponseDto(false, err.Error())
		}
		return *binding.NewResponseDto(true, "Successful")
	}

	return *binding.NewResponseDto(false, "Incorrect password supplied")

}
