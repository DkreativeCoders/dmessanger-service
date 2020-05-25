package domain

//Customer Entity containing basic fields and userDetails
//swagger:model customer-model
type Customer struct {
	User                    `gorm:"foreignkey:UserId"`
	UserId                 uint
	DefaultShippingAddress string
	TotalNumberOfOrders    int `gorm:"default:0"`
}
