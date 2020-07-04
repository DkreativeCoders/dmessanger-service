package service

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/config/mail"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/mocks"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCustomerService_CreateUser(t *testing.T) {

	//	mock userRepository
	//	mock customerRepository
	//	mock tokenRepository
	//	mock mailService

	if testing.Short() {
		t.Skip()
	}

	otpNumber := "000111"
	timeAdded := time.Now().Add(1 * time.Hour)
	mailtobesent := mail.NewEMailMessage("DkreativeCoders Verify User",
		"Please visit this link to verify your account. \n This links expires in an hour \n"+"https://dmessanger-service.herokuapp.com/verify-user/unique-111\n You can also use this OTP to verify your account via your mobile Device "+otpNumber, "daniel@gmail.com",
		nil)

	testCases := []struct {
		name    string
		request dto.CustomerRequest

		userRepoFindUserExistInput      string
		userRepoFindUserExistReturnData bool

		customerRepositorySaveInput        domain.Customer
		customerRepositorySaveReturnOutput *domain.Customer
		customerRepositorySaveReturnError  error

		uniqueIdGenerated string

		tokenServiceInputUserID               uint
		tokenServiceInputUniqueID             string
		tokenServiceInputExpirationTimeInHour time.Duration
		tokenServiceCreateReturnOutPut        *domain.Token
		tokenServiceCreateReturnError         error

		mailServiceSendMailInput  mail.EMailMessage
		mailServiceSendMailOutput string
		mailServiceSendMailError  error

		expectedValueDataResponse  *domain.Customer
		expectedValueErrorResponse error
	}{
		{
			"Test with valid customer request",
			dto.CustomerRequest{FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"},

			"daniel@gmail.com",
			false,

			domain.Customer{User: domain.User{FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"}},
			&domain.Customer{User: domain.User{Model: gorm.Model{ID: 0}, FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"}},
			nil,

			"unique-111",

			0,
			"unique-111",
			1,
			&domain.Token{Model: gorm.Model{ID: 0}, UserId: 0, Token: "unique-111", ExpiresOn: timeAdded},
			nil,

			*mailtobesent,
			"success",
			nil,

			&domain.Customer{User: domain.User{Model: gorm.Model{ID: 0}, FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"}},
			nil,
		},

		{
			"Test customer already exist with already registered email",
			dto.CustomerRequest{FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"},

			"daniel@gmail.com",
			true,

			domain.Customer{User: domain.User{FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"}},
			&domain.Customer{User: domain.User{Model: gorm.Model{ID: 0}, FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"}},
			nil,

			"unique-111",

			0,
			"unique-111",
			1,
			&domain.Token{Model: gorm.Model{ID: 0}, UserId: 0, Token: "unique-111", ExpiresOn: timeAdded},
			nil,

			*mailtobesent,
			"success",
			nil,

			nil,
			errors.New("user Already Exist with email"),
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			// Create dependency userRepo with mock implementation
			userRepo := mocks.IUserRepository{}
			userRepo.On("FindUserExist", testCase.userRepoFindUserExistInput).Return(testCase.userRepoFindUserExistReturnData)

			// Create dependency customerRepo with mock implementation
			customerRepo := mocks.ICustomerRepository{}
			customerRepo.On("Save", testCase.customerRepositorySaveInput).Return(testCase.customerRepositorySaveReturnOutput, testCase.customerRepositorySaveReturnError)

			// Create dependency tokenRepo with mock implementation
			tokenService := mocks.ITokenService{}
			tokenService.On("CreateTokenWithExpirationInHours", testCase.tokenServiceInputUserID, testCase.tokenServiceInputUniqueID, testCase.tokenServiceInputExpirationTimeInHour).Return(testCase.tokenServiceCreateReturnOutPut, testCase.tokenServiceCreateReturnError)

			tokenRepo := mocks.ITokenRepository{}

			// Create dependency tokenRepo with mock implementation
			uuidService := mocks.IUuid{}
			uuidService.On("GenerateUniqueId").Return(testCase.uniqueIdGenerated)

			// Create dependency tokenRepo with mock implementation
			mailService := mocks.IMail{}
			mailService.On("SendEMail", testCase.mailServiceSendMailInput).Return(testCase.mailServiceSendMailOutput, testCase.mailServiceSendMailError)

			//Todo: OTP service to be replaced with mock
			otpService := mocks.IOtp{}
			otpService.On("GenerateOTP").Return(otpNumber)

			// Create userService and inject mock repo
			customerService := INewCustomerService(&customerRepo, &userRepo, &tokenRepo, &tokenService, &mailService, &uuidService, &otpService)

			// Actual method call
			output, err := customerService.CreateUser(testCase.request)

			if err != nil {
				assert.Equal(t, testCase.expectedValueErrorResponse, err)
			}
			// Expected output
			expected := testCase.expectedValueDataResponse

			assert.Equal(t, expected, output)
		})
	}

}
