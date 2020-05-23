package domain

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	UserId                 uint
	User                   User `gorm:"foreignkey:UserId"`
	DefaultShippingAddress string
	TotalNumberOfOrders    int `gorm:"default:0"`
}
