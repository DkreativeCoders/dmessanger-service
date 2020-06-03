package domain

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Token struct {
	gorm.Model
	UserId    uint
	Token     string
	ExpiresOn time.Time
}
