package domain

import "github.com/jinzhu/gorm"

//Customer Entity containing basic fields and userDetails
//swagger:model courierModel
type Courier struct {
	gorm.Model
	UserId                 uint
	User                   User `gorm:"foreignkey:UserId"`
	Rating                 int
	NumberOfRides          int
	NumberOfCompletedRides int
}
