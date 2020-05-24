package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}



//swagger:model responseModel
type ResponseDto struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
}

func NewResponseDto(status bool, message string) *ResponseDto {
	var response *ResponseDto
	response.Message = message
	response.Status = status
	return response
}

