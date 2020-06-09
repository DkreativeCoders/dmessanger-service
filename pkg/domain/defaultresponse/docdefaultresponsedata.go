package defaultresponse

// A Default Response Message
//swagger:response responseDto
type ResponseWrapper struct {
	// The Response message
	// in: body
	Body ResponseData
}

// A  Bad Request Error Response .
//swagger:response badRequestResponse
type BadRequestResponseWrapper struct {
	// in: body
	Body ResponseData
}

// A UnAuthenticated Error Response .
//swagger:response unAuthenticatedResponse
type UnAuthenticatedResponseWrapper struct {
	// in: body
	Body ResponseData
}