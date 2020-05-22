package domain_test

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/domain"
	"reflect"
	"testing"
)

func TestUser_Validate(t *testing.T) {

	testCases := []struct {
		name           string
		user           domain.User
		expected       map[string]interface{}
		expectedStatus bool
	}{
		{ "Test with empty first name",
			domain.User{FirstName: "", LastName: "Mark", Age: "34", Email: "adam@gmail.com", PhoneNumber: "01-7463-546", Password: "password", Address: "12 Akobi Crescent"},
			map[string]interface{}{"status": false, "message": "User First name should be on the payload"},
			false,
		},
		{ "Test with empty last name",
			domain.User{FirstName: "Adam", LastName: "", Age: "34", Email: "adam@gmail.com", PhoneNumber: "01-7463-546", Password: "password", Address: "12 Akobi Crescent"},
			map[string]interface{}{"status": false, "message": "User Last name should be on the payload"},
			false,
		},
		{ "Test with empty age",
			domain.User{FirstName: "Adam", LastName: "Mark", Age: "", Email: "adam@gmail.com", PhoneNumber: "01-7463-546", Password: "password", Address: "12 Akobi Crescent"},
			map[string]interface{}{"status": false, "message": "User Age should be on the payload"},
			false,
		},
		{ "Test with empty phone number",
			domain.User{FirstName: "Adam", LastName: "Mark", Age: "34", Email: "adam@gmail.com", PhoneNumber: "", Password: "password", Address: "12 Akobi Crescent"},
			map[string]interface{}{"status": false, "message": "User Phone number should be on the payload"},
			false,
		},
		{ "Test with empty email",
			domain.User{FirstName: "Adam", LastName: "Mark", Age: "34", Email: "", PhoneNumber: "01-7463-546", Password: "password", Address: "12 Akobi Crescent"},
			map[string]interface{}{"status": false, "message": "User Email should be on the payload"},
			false,
		},
		{ "Test with valid arguments",
			domain.User{FirstName: "Adam", LastName: "Mark", Age: "34", Email: "adam@gmail.com", PhoneNumber: "01-7463-546", Password: "password", Address: "12 Akobi Crescent"},
			map[string]interface{}{"status": true, "message": "success"},
			true,
		},
	}

	for _, testCase := range testCases{
		t.Run(testCase.name, func(t *testing.T) {
			// run user validate function
			output, outputStatus := testCase.user.Validate()

			// check output status
			if outputStatus != testCase.expectedStatus {
				t.Fail()
			}

			// check output
			if !reflect.DeepEqual(output, testCase.expected) {
				t.FailNow()
			}
		})
	}
}
