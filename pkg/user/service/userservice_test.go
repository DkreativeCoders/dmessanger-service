package service_test

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/mocks"
	"github.com/DkreativeCoders/dmessanger-service/pkg/user/service"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CreateUser(t *testing.T) {

	testCases := []struct{
		name string
		repoInputData domain.User
		repoReturnData *domain.User
		repoReturnErr error
		expectedVal map[string]interface{}
	}{
		{
			"Test with valid user input",
			domain.User{FirstName:"Adam", LastName:"Mark", Age:"24", Email:"amark@gmail.com", PhoneNumber:"01-2345-6789", Password:"password", Address: "401, Hebert Mark Way"},
			&domain.User{Model: gorm.Model{ID:0}, FirstName:"Adam", LastName:"Mark", Age:"24", Email:"amark@gmail.com", PhoneNumber:"01-2345-6789", Password:"password", Address: "401, Hebert Mark Way"},
			nil,
			map[string]interface{}{
				"status": true,
				"message": "success",
				"data": &domain.User{Model: gorm.Model{ID:0}, FirstName:"Adam", LastName:"Mark", Age:"24", Email:"amark@gmail.com", PhoneNumber:"01-2345-6789", Password:"password", Address: "401, Hebert Mark Way"},
			},
		},
		{
			"Test with error from repository",
			domain.User{Model: gorm.Model{ID:0}, FirstName:"Adam", LastName:"Mark", Age:"24", Email:"amark@gmail.com", PhoneNumber:"01-2345-6789", Password:"password", Address: "401, Hebert Mark Way"},
			nil,
			errors.New("user already exist"),
			map[string]interface{}{
				"status": false,
				"message": "Error",
				"error_message": errors.New("user already exist"),
			},
		},
		{
			"Test with empty first name",
			domain.User{FirstName:"", LastName:"Mark", Age:"24", Email:"amark@gmail.com", PhoneNumber:"01-2345-6789", Password:"password", Address: "401, Hebert Mark Way"},
			&domain.User{Model: gorm.Model{ID:0}, FirstName:"", LastName:"Mark", Age:"24", Email:"amark@gmail.com", PhoneNumber:"01-2345-6789", Password:"password", Address: "401, Hebert Mark Way"},
			nil,
			map[string]interface{}{
				"status": false,
				"message": "User First name should be on the payload",
			},
		},
	}

	for _, testCase := range testCases{

		t.Run(testCase.name, func(t *testing.T) {
			// Create dependency userRepo with mock implementation
			userRepo := mocks.UserRepository{}
			userRepo.On("Save", testCase.repoInputData).Return(testCase.repoReturnData, testCase.repoReturnErr)

			// Create userService and inject mock repo
			userService := service.INewService(&userRepo)

			// Actual method call
			output := userService.CreateUser(testCase.repoInputData)

			// Expected output
			expected := testCase.expectedVal

			assert.Equal(t, expected, output)
		})
	}
}

//: Currently failing cos GetUser has not been implemented
//Todo: add more test cases
func TestService_GetUser(t *testing.T) {

	if testing.Short(){
		t.Skip("Skipped test: The test fails cos it hasn't been implemented")
	}

	testCases := []struct{
		name string
		repoInputData int
		repoReturnData *domain.User
		repoReturnErr error
		expectedVal *domain.User
		expectedErr error
	}{
		{
			"Test with valid user input",
			101,
			&domain.User{Model: gorm.Model{ID:0}, FirstName:"Adam", LastName:"Mark", Age:"24", Email:"amark@gmail.com", PhoneNumber:"01-2345-6789", Password:"password", Address: "401, Hebert Mark Way"},
			nil,
			&domain.User{Model: gorm.Model{ID:0}, FirstName:"Adam", LastName:"Mark", Age:"24", Email:"amark@gmail.com", PhoneNumber:"01-2345-6789", Password:"password", Address: "401, Hebert Mark Way"},
			nil,
		},
	}

	for _, testCase := range testCases{

		t.Run(testCase.name, func(t *testing.T) {
			// Create dependency userRepo with mock implementation
			userRepo := mocks.UserRepository{}
			userRepo.On("FindByID", testCase.repoInputData).Return(testCase.repoReturnData, testCase.repoReturnErr)

			// Create userService and inject mock repo
			userService := service.INewService(&userRepo)

			// Actual method call
			output, err := userService.GetUser(testCase.repoInputData)
			if err != nil{
				assert.Equal(t, testCase.expectedVal, err)
			}
			assert.Equal(t, testCase.expectedVal, output)
		})
	}
}

func TestService_GetAllUser(t *testing.T) {
	testCases := []struct{
		name string
		repoReturnData []domain.User
		expectedVal map[string]interface{}
	}{
		{
			"Test with numerous user in storage",
			[]domain.User{
				domain.User{Model: gorm.Model{ID:0}, FirstName:"Adam", LastName:"Mark", Age:"24", Email:"amark@gmail.com", PhoneNumber:"01-2345-6789", Password:"password", Address: "401, Hebert Mark Way"},
				domain.User{Model: gorm.Model{ID:1}, FirstName:"Tony", LastName:"Young", Age:"37", Email:"tonyyoung@gmail.com", PhoneNumber:"01-5678-6789", Password:"youngTony", Address: "78, Broad Street"},
				domain.User{Model: gorm.Model{ID:2}, FirstName:"Ross", LastName:"Barkley", Age:"50", Email:"barks@gmail.com", PhoneNumber:"01-3333-8907", Password:"BarksRSS", Address: "46B, Moore Street"},
				},
			map[string]interface{}{
				"status": true,
				"message": "success",
				"data": []domain.User{
					domain.User{Model: gorm.Model{ID:0}, FirstName:"Adam", LastName:"Mark", Age:"24", Email:"amark@gmail.com", PhoneNumber:"01-2345-6789", Password:"password", Address: "401, Hebert Mark Way"},
					domain.User{Model: gorm.Model{ID:1}, FirstName:"Tony", LastName:"Young", Age:"37", Email:"tonyyoung@gmail.com", PhoneNumber:"01-5678-6789", Password:"youngTony", Address: "78, Broad Street"},
					domain.User{Model: gorm.Model{ID:2}, FirstName:"Ross", LastName:"Barkley", Age:"50", Email:"barks@gmail.com", PhoneNumber:"01-3333-8907", Password:"BarksRSS", Address: "46B, Moore Street"},
				},
			},
		},
		{
			"Test with no user in storage",
			[]domain.User{},
			map[string]interface{}{
				"status": true,
				"message": "success",
				"data": []domain.User{},
			},
		},
	}

	for _, testCase :=range testCases{
		userRepo := mocks.UserRepository{}
		userRepo.On("FindAll").Return(testCase.repoReturnData)

		// Create userService and inject mock repo
		userService := service.INewService(&userRepo)

		// Actual method call
		output := userService.GetAllUser()

		// Assertion
		assert.Equal(t, testCase.expectedVal, output)
	}
}

func TestService_UpdatePassword(t *testing.T) {

}