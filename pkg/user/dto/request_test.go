package dto

import (
	"errors"
	"github.com/DkreativeCoders/dmessanger-service/pkg/customer/dto"
	"reflect"
	"testing"
)

func TestUser_Validate(t *testing.T) {

	testCases := []struct {
		name            string
		customerRequest dto.CustomerRequest
		expected        error
	}{
		{"Test with empty first name",
			dto.CustomerRequest{FirstName: "", LastName: "Mark", Age: "34", Email: "adam@gmail.com", PhoneNumber: "01-7463-546", Password: "password", Address: "12 Akobi Crescent"},
			errors.New("firstName cannot be empty"),
		},
		{"Test with empty last name",
			dto.CustomerRequest{FirstName: "Adam", LastName: "", Age: "34", Email: "adam@gmail.com", PhoneNumber: "01-7463-546", Password: "password", Address: "12 Akobi Crescent"},
			errors.New("lastName cannot be empty"),
		},
		{"Test with empty age",
			dto.CustomerRequest{FirstName: "Adam", LastName: "Mark", Age: "", Email: "adam@gmail.com", PhoneNumber: "01-7463-546", Password: "password", Address: "12 Akobi Crescent"},
			errors.New("age  cannot be empty"),
		},
		{"Test with empty phone number",
			dto.CustomerRequest{FirstName: "Adam", LastName: "Mark", Age: "34", Email: "adam@gmail.com", PhoneNumber: "", Password: "password", Address: "12 Akobi Crescent"},
			errors.New("phone number cannot be empty"),
		},
		{"Test with empty email",
			dto.CustomerRequest{FirstName: "Adam", LastName: "Mark", Age: "34", Email: "", PhoneNumber: "01-7463-546", Password: "password", Address: "12 Akobi Crescent"},
			errors.New("email cannot be empty"),
		},
		{"Test with empty password",
			dto.CustomerRequest{FirstName: "Adam", LastName: "Mark", Age: "34", Email: "email@email.com", PhoneNumber: "01-7463-546", Password: "", Address: "12 Akobi Crescent"},
			errors.New("password cannot be empty"),
		},
		{"Test with valid arguments",
			dto.CustomerRequest{FirstName: "Adam", LastName: "Mark", Age: "34", Email: "adam@gmail.com", PhoneNumber: "01-7463-546", Password: "password", Address: "12 Akobi Crescent"},
			nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// run user validate function
			err := testCase.customerRequest.Validate()

			// check output
			if !reflect.DeepEqual(err, testCase.expected) {
				t.FailNow()
			}
		})
	}
}
