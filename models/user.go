package models

import (
	"github.com/danieloluwadare/dmessanger/utils"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName  string
	LastName  string
	Age string
	Email string
	PhoneNumber string
	Address string
}

func (user *User) Validate() (map[string]interface{}, bool) {

	if user.FirstName == "" {
		return utils.Message(false, "User First name should be on the payload"), false
	}

	if user.LastName == "" {
		return utils.Message(false, "User Last name should be on the payload"), false
	}

	if user.PhoneNumber == "" {
		return utils.Message(false, "User Phone number should be on the payload"), false
	}

	if user.Age == "" {
		return utils.Message(false, "User Age should be on the payload"), false
	}

	if user.Email == "" {
		return utils.Message(false, "User Email should be on the payload"), false
	}

	//All the required parameters are present
	return utils.Message(true, "success"), true
}
