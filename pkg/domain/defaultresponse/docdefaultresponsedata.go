package defaultresponse
// A Default is the default Response.
//swagger:response responseDto
type ResponseWrapper struct {
	// The Response message
	// in: body
	Body struct {
		ResponseData
	}
}