package uuid

import (
	uuid "github.com/satori/go.uuid"
)

func INewUuid() IUuid {
	return uuidService{}
}

type uuidService struct {
}

func (u uuidService) GenerateUniqueId() string {
	uniqueId := uuid.NewV4().String()
	return uniqueId
}
