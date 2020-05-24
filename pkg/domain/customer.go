package domain

import "github.com/jinzhu/gorm"

//Customer Entity containing basic fields and userDetails
//swagger:model customerModel
type Customer struct {
	gorm.Model
	UserId                 uint
	User                    `gorm:"foreignkey:UserId"`
	DefaultShippingAddress string
	TotalNumberOfOrders    int `gorm:"default:0"`
}
