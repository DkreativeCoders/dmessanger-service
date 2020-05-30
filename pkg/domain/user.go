package domain

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/utils"
	"github.com/jinzhu/gorm"
)

//User Entity  containing basic fields
//swagger:model user-model
type User struct {
	gorm.Model
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Age         string `json:"age"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"-"`
	Address     string `json:"address"`
	IsEnabled   bool   `json:"isEnabled"`
	IsActive  bool   `json:"isActive"`
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

func (user *User) ValidateToError() error {

	if user.FirstName == "" {
		return errors.New("user first name should be on the payload")
	}

	if user.LastName == "" {
		return errors.New("user last name should be on the payload")

	}

	if user.PhoneNumber == "" {
		return errors.New("user phone number should be on the payload")
	}

	if user.Age == "" {
		return errors.New("user age  should be on the payload")

	}

	if user.Email == "" {
		return errors.New("user email should be on the payload")

	}

	//All the required parameters are present
	return nil
}
