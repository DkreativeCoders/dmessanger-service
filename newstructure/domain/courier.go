package domain

import "github.com/jinzhu/gorm"

type Courier struct {
	gorm.Model
	UserId uint
	User User  `gorm:"foreignkey:UserId"`
	Rating int
	NumberOfRides int
	NumberOfCompletedRides int
}
