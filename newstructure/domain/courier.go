package domain

import "github.com/jinzhu/gorm"

type Courier struct {
	gorm.Model
	UserID uint
	User User  `gorm:"foreignkey:UserID"`
	Rating int
	NumberOfRides int
	NumberOfCompletedRides int
}
