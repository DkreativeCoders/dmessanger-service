package domain

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	UserId                 uint
	User                   User `gorm:"foreignkey:UserId"`
	defaultShippingAddress string
	totalNumberOfOrders         int `gorm:"default:0"`
}
