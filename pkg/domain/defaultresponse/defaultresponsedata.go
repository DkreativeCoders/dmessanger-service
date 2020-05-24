package defaultresponse

//swagger:model defaultResponseModel
type ResponseData struct {
	// The ResponseData status
	//
	// Required: true
	Status  bool        `json:"status"`
	// Required: true
	Message string      `json:"message"`
}

func NewResponseDto(status bool, message string) *ResponseData {
	response := ResponseData{}
	response.Status = status
	response.Message = message
	return &response
}
