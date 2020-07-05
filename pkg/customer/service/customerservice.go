package service

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/mail"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/otp"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/uuid"
	"github.com/DkreativeCoders/dmessanger-service/pkg/constanst"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/irepository"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain/iservice"
	"os"
	"time"
)

//INewService return an interface that's why Constrictor/Method name is preceded with I
func INewCustomerService(
	repository irepository.ICustomerRepository,
	userRepository irepository.IUserRepository,
	tokenRepository irepository.ITokenRepository,
	tokenService iservice.ITokenService,
	mailService mail.IMail,
	uuid uuid.IUuid,
	otp otp.IOtp,
) iservice.ICustomerService {
	return customerService{
		repository,
		userRepository,
		tokenRepository,
		tokenService,
		mailService,
		uuid,
		otp,
	}
}

type customerService struct {
	customerRepository irepository.ICustomerRepository
	userRepository     irepository.IUserRepository
	tokenRepository    irepository.ITokenRepository
	tokenService       iservice.ITokenService
	mailService        mail.IMail
	uuid               uuid.IUuid
	otp                otp.IOtp
}

//refactor and test case needed
//Validate and crease customer
func (s customerService) CreateUser(request dto.CustomerRequest) (*domain.Customer, error) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	if found := s.userRepository.FindUserExist(request.Email); found {
		return nil, errors.New("user Already Exist with email")
	}

	customer := domain.Customer{}
	customer.FirstName = request.FirstName
	customer.LastName = request.LastName
	customer.Email = request.Email
	customer.Address = request.Address
	customer.PhoneNumber = request.PhoneNumber
	customer.Password = request.Password
	customer.Age = request.Age

	newCustomer, err := s.customerRepository.Save(customer)
	if err != nil {
		return nil, err
	}

	return newCustomer, nil
}

func (s customerService) SendActivationMail(customer domain.Customer) error {
	// Generate unique id and otp
	uniqueId := s.uuid.GenerateUniqueId()
	otp := s.otp.GenerateOTP()

	// Create unique token for customer
	_, err := s.tokenService.CreateTokenWithExpirationInHours(customer.UserId, uniqueId, otp, 1)
	if err != nil {
		return err
	}

	// Create customer verification link
	var linkToSend string
	hostName := os.Getenv("HOST_NAME")
	if hostName == ""{
		linkToSend = "https://dmessanger-service.herokuapp.com/api/v1/customers/verify?token=" + uniqueId
	}else{
		linkToSend = hostName + constanst.ApiVersion1 + "customers/verify?token=" + uniqueId
	}

	// Create email
	email := s.createMail(customer, linkToSend, otp)

	// Send activation mail
	_, er := s.mailService.SendEMail(*email)
	if er != nil {
		return errors.New("error occurred while trying to send mail, try again later")
	}
	return nil
}

func (s customerService) ActivateUser(tk string) error {
	// Check for token in the repository
	token, err := s.tokenRepository.FindByToken(tk)
	if err != nil {
		return err
	}

	// Check if token has expired
	tokenExpiryTime := token.ExpiresOn
	if time.Now().After(tokenExpiryTime) {
		return errors.New("token expired")
	}

	// Find user with token
	user, er := s.userRepository.FindByID(int(token.UserId))
	if er != nil {
		return er
	}

	// Activate user
	user.IsActive = true
	if _, err := s.userRepository.Update(*user); err != nil{
		return err
	}

	return nil
}

func (s customerService) createMail(customer domain.Customer, linkToSend, otp string) *mail.EMailMessage {
	subject := "DMessanger User Verification"
	text := "Please visit this link to verify your account. \n This links expires in an hour \n" + linkToSend + "." +
		"\n You can also use this OTP to verify your account via your mobile Device " + otp
	recipient := customer.Email
	email := mail.NewEMailMessage(subject, text, recipient, nil)
	return email
}