package binding

//swagger:model responseData
type ResponseData struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// swagger:response responseDto
type ResponseDto struct {
	ResponseData
}

func NewResponseDto(status bool, message string) *ResponseDto {
	response := ResponseDto{}
	response.Status = status
	response.Message = message
	return &response
}

