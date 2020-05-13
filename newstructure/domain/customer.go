package domain

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	UserID uint
	User User  `gorm:"foreignkey:UserID"`
}
