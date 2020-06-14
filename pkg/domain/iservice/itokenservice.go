package iservice

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"time"
)

type ITokenService interface {
	CreateTokenWithExpirationInHours(userID uint, uniqueID string, expirationTimeInHour time.Duration) (*domain.Token, error)
	CreateTokenWithExpirationInMinutes(userID uint, uniqueID string, expirationTimeInMinutes time.Duration) (*domain.Token, error)
}
