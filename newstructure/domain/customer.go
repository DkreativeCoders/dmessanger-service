package domain

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	User User
}
