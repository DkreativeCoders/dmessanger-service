package utils_test

import (
	"encoding/json"
	"github.com/DkreativeCoders/dmessanger-service/pkg/utils"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMessage(t *testing.T) {
	testCases := []struct {
		name    string
		status  bool
		message string
		output  map[string]interface{}
	}{
		{"Test that status and message yield required value", true, "Operation Successful", map[string]interface{}{"status": true, "message": "Operation Successful"}},
		{"Test that status and empty message yields required value", false, "", map[string]interface{}{"status": false, "message": ""}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			output := utils.Message(testCase.status, testCase.message)
			if !reflect.DeepEqual(output, testCase.output) {
				t.Fail()
			}
		})
	}
}

func TestRespond(t *testing.T) {
	// Test inputs
	w := httptest.NewRecorder() // records the operation output given as a response
	data := map[string]interface{}{"status": true, "message": "Operation Successful"}

	// Test function
	utils.Respond(w, data)

	// Retrieves response and decodes the data
	response := w.Result()
	output := make(map[string]interface{})
	err := json.NewDecoder(response.Body).Decode(&output)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Compare input data to output
	if !reflect.DeepEqual(output, data) {
		t.FailNow()
	}
}
