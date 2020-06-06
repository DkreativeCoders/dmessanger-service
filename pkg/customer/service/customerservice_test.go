package service

import (
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

	timeAdded := time.Now().Add(1 * time.Hour)

	testCases := []struct {
		name           string
		request  dto.CustomerRequest

		userRepoFindUserExistInput string
		userRepoFindUserExistReturnData bool

		customerRepositorySaveInput domain.Customer
		customerRepositorySaveReturnOutput *domain.Customer
		customerRepositorySaveReturnError error

		uniqueIdGenerated string


		tokenRepositoryCreateInput domain.Token
		tokenRepositoryCreateReturnOutPut *domain.Token
		tokenRepositoryCreateReturnError error


		mailServiceSendMailInput *mail.EMailMessage
		mailServiceSendMailOutput string
		mailServiceSendMailError error

		expectedValueDataResponse *domain.Customer
		expectedValueErrorResponse error

	}{
		{
			"Test with valid customer request",
			dto.CustomerRequest{FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"},

			"daniel@gmail.com",
			false,

			domain.Customer{User:domain.User{FirstName: "Daniel", LastName: "Dada", Age: "20", Email: "daniel@gmail.com", PhoneNumber: "08282888", Password: "password", Address: "address daniel"},},
			&domain.Customer{User:domain.User{Model: gorm.Model{ID: 0},FirstName:"Daniel",LastName:"Dada",Age:"20",Email:"daniel@gmail.com", PhoneNumber:"08282888",Password:"password",Address:"address daniel"}},
			nil,

			"unique-111",

			domain.Token{UserId: 0,Token: "unique-111",ExpiresOn: timeAdded},
			&domain.Token{Model: gorm.Model{ID: 0},UserId: 0,Token: "random string",ExpiresOn: timeAdded},
			nil,

			mail.NewEMailMessage("DkreativeCoders Verify User",
				"Please visit this link to verify your account. \n This links expires in an hour \n" + "unique-111","daniel@gmail.com",
				nil),
			"success",
			nil,

			&domain.Customer{User:domain.User{Model: gorm.Model{ID: 0},FirstName:"Daniel",LastName:"Dada",Age:"20",Email:"daniel@gmail.com", PhoneNumber:"08282888",Password:"password",Address:"address daniel"}},
nil,
		},
	}


	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			// Create dependency userRepo with mock implementation
			userRepo := mocks.IUserRepository{}
			userRepo.On("FindUserExist", testCase.userRepoFindUserExistInput).Return(testCase.userRepoFindUserExistReturnData)

			// Create dependency customerRepo with mock implementation
			customerRepo := mocks.ICustomerRepository{}
			customerRepo.On("Save", testCase.customerRepositorySaveInput).Return(testCase.customerRepositorySaveReturnOutput,testCase.customerRepositorySaveReturnError)

			// Create dependency tokenRepo with mock implementation
			tokenRepo := mocks.ITokenRepository{}
			tokenRepo.On("Create", testCase.tokenRepositoryCreateInput).Return(testCase.tokenRepositoryCreateReturnOutPut,testCase.tokenRepositoryCreateReturnError)

			// Create dependency tokenRepo with mock implementation
			uuidService := mocks.IUuid{}
			uuidService.On("GenerateUniqueId").Return(testCase.uniqueIdGenerated)

			// Create dependency tokenRepo with mock implementation
			mailService := mocks.IMail{}
			mailService.On("SendEMail", testCase.mailServiceSendMailInput).Return(testCase.mailServiceSendMailOutput,testCase.mailServiceSendMailError)



			// Create userService and inject mock repo
			customerService := INewCustomerService(&customerRepo,&userRepo,&tokenRepo,&mailService,&uuidService)

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