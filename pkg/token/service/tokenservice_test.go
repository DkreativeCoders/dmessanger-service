package service

import (
	"bou.ke/monkey"
	otp2 "github.com/DkreativeCoders/dmessanger-service/pkg/config/otp"
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"github.com/DkreativeCoders/dmessanger-service/pkg/mocks"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTokenService_CreateTokenWithExpirationInHours(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	timeTestInputForRepo := time.Date(2020, 06, 17, 20, 34, 58, 651387237, time.UTC)
	timeTest2 := timeTestInputForRepo.Add(1 * time.Hour)

	//timeTest2 := time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return timeTestInputForRepo
	})

	testCases := []struct {
		name                 string
		UserID               uint
		uniqueID             string
		expirationTimeInHour time.Duration
		repoTokenInputData   domain.Token
		repoTokenReturnData  *domain.Token
		repoReturnErr        error
		expectedVal          *domain.Token
		expectedErr          error
	}{
		{
			"Test with valid  input",
			1,
			"unique-id",
			1,
			domain.Token{UserId: 1, Token: "unique-id", ExpiresOn: timeTest2},
			&domain.Token{Model: gorm.Model{ID: 0}, UserId: 1, Token: "unique-id", ExpiresOn: timeTest2},
			nil,
			&domain.Token{Model: gorm.Model{ID: 0}, UserId: 1, Token: "unique-id", ExpiresOn: timeTestInputForRepo.Add(1 * time.Hour)},
			nil,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			// Create dependency userRepo with mock implementation
			tokenRepo := mocks.ITokenRepository{}
			tokenRepo.On("Create", testCase.repoTokenInputData).Return(testCase.repoTokenReturnData, testCase.repoReturnErr)

			// Create userService and inject mock repo
			tokenService := INewTokenService(&tokenRepo)

			//Todo: OTP service to be replaced with mock
			otp := otp2.NewOTPService()

			// Actual method call
			output, _ := tokenService.CreateTokenWithExpirationInHours(testCase.UserID, testCase.uniqueID, otp.GenerateOTP(), testCase.expirationTimeInHour)

			monkey.Unpatch(time.Now)
			// Expected output
			expected := testCase.expectedVal

			assert.Equal(t, expected, output)
		})
	}
}
