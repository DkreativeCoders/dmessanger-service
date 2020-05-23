package pkg

import (
	"github.com/DkreativeCoders/dmessanger-service/pkg/utils"
)


//swagger:response responseDto
type ResponseWrapper struct {
	// in: body
	Body struct {
		utils.ResponseDto
	}
}