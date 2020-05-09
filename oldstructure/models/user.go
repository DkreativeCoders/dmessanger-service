package models

import (
	"github.com/danieloluwadare/dmessanger/oldstructure/utils"
	"github.com/jinzhu/gorm"
)

//User Entity  containing basic fields
type User struct {
	gorm.Model
	FirstName  string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age string `json:"age"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address string `json:"address"`
}

//Validate All the fields for User
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
