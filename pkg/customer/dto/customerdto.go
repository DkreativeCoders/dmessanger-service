package dto

import (
	"errors"
)

//swagger:model
type CustomerRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Age         string `json:"age"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"-"`
	Address     string `json:"address"`
	//defaultShippingAddress 		string
	//totalNumberOfOrders         string
}


//Validate all the fields of the customer
func (request *CustomerRequest) ValidateAll() (bool, []error){

	foundError := false
	arrErrors := make([]error, 0)

	if request.FirstName == "" {
		foundError=true
		arrErrors = append(arrErrors, errors.New("firstName cannot be empty"))
	}

	if request.LastName == "" {
		foundError=true
		arrErrors = append(arrErrors, errors.New("lastName cannot be empty"))
		//return strings.New("LastName cannot be empty")
	}

	if request.PhoneNumber == "" {
		foundError=true
		arrErrors = append(arrErrors, errors.New("phone number cannot be empty"))
	}

	if request.Age == "" {
		foundError=true
		arrErrors = append(arrErrors, errors.New("age cannot be empty"))
	}

	if request.Email == "" {
		foundError=true
		arrErrors = append(arrErrors, errors.New("email cannot be empty"))
	}

	if foundError{
		return foundError,arrErrors
	}

	return foundError,nil

	//All the required parameters are present
}

func (request *CustomerRequest) Validate() error{

	if request.FirstName == "" {
		return errors.New("firstName cannot be empty")
	}

	if request.LastName == "" {
		return errors.New("lastName cannot be empty")
	}

	if request.PhoneNumber == "" {
		return errors.New("phone number cannot be empty")
	}

	if request.Age == "" {
		return errors.New("age  cannot be empty")
	}

	if request.Email == "" {
		return errors.New("email cannot be empty")
	}

	//All the required parameters are present
	return nil

}