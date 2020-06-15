package iservice

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"time"
)

type ITokenService interface {
	CreateTokenWithExpirationInHours(userID uint, uniqueID, otp string, expirationTimeInHour time.Duration) (*domain.Token, error)
	CreateTokenWithExpirationInMinutes(userID uint, uniqueID, otp string, expirationTimeInMinutes time.Duration) (*domain.Token, error)
}
