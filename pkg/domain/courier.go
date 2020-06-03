package domain

import "github.com/jinzhu/gorm"

//Courier Entity containing basic fields and userDetails
//swagger:model courier-model
type Courier struct {
	gorm.Model
	UserId                 uint
	User                   User `gorm:"foreignkey:UserId"`
	Rating                 int
	NumberOfRides          int
	NumberOfCompletedRides int
}
