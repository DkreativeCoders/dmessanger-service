package defaultresponse
type ResponseData struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponseDto(status bool, message string, data interface{}) *ResponseData {
	response := ResponseData{}
	response.Status = status
	response.Message = message
	response.Data=data
	return &response
}
