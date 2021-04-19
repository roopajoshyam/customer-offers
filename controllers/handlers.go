package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
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

	user := &models.User{}
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
	log.Println("user Id", user.Email)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdUser)
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := FindUser(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}

func FindUser(email, password string) map[string]interface{} {
	user := &models.User{}
	err := db.Where("Email = ?", email).First(user).Error
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email id not present"}
		return resp
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials"}
		return resp
	}

	tk := &models.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}

func FetchUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db.First(&user, id)
	json.NewEncoder(w).Encode(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db.First(&user, id)
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db.Delete(&user, id)
	json.NewEncoder(w).Encode(&user)
}


