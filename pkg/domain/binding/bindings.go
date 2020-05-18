package binding

//swagger:model responseData
type responseData struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

// swagger:response responseDto
type ResponseDto struct {
	responseData
}


