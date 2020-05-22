package domain

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	userId                 uint
	User                   User `gorm:"foreignkey:UserId"`
	defaultShippingAddress string
	totalNumberOfOrders         string
}
