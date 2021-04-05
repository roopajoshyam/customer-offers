package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jroopa/offers/pkg/models"
	"github.com/jroopa/offers/utils"

	"golang.org/x/crypto/bcrypt"
)

var db = utils.ConnectDB()

type ErrorResponse struct {
	Err string
}

func TestAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API live and kicking"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{
		Name:  "roopa",
		Email: "jroopa@gmail.com",
	}
	json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)
	createdUser := db.Debug().Create(user)
	log.Println("user Id", user.ID)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdUser)
}
