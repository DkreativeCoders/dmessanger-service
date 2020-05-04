package main

import (
	// 	"log"
	// 	"net/http"

	// 	"github.com/gorilla/mux"
	"fmt"

	"github.com/danieloluwadare/dmessanger/models"
	"github.com/danieloluwadare/dmessanger/interfaces"

)

type Tetse struct {
	id int
}

func (u Tetse) GetUser(id int) *models.User {
	fmt.Println("ooo")
	return nil
}
func (u Teste) CreateUser(user models.User) error {
	fmt.Println("ooo")
	return nil

}

func main() {
	var face interfaces.IUserRepository  
	 test := Tetse{1};
	 test.
}
