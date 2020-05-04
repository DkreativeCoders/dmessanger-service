package controllers

import (
	"net/http"

	"github.com/danieloluwadare/dmessanger/interfaces/iservice"
	"github.com/danieloluwadare/dmessanger/models"
	"github.com/danieloluwadare/dmessanger/repository"
	"github.com/danieloluwadare/dmessanger/service"
)

var UserController UserHandler

func init() {

	db := models.GetDB()
	repository := repository.NewInMemoryRepository(db)
	service := service.NewService(repository)
	handler := NewBookHandler(service)

	UserController = handler

}

func NewBookHandler(userService iservice.IUserService) UserHandler {
	return UserHandler{userService}
}

type UserHandler struct {
	userService iservice.IUserService
}

func (u UserHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
}

func (u UserHandler) GetBook(w http.ResponseWriter, r *http.Request) {

}
