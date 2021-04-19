package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jroopa/offers/pkg/models"
)

func CreateOffer(w http.ResponseWriter, r *http.Request) {
	offer := &models.Offers{}
	json.NewDecoder(r.Body).Decode(offer)
	log.Printf(" %v",offer.UserID)
	var user models.User
	userInfo := db.Debug().First(&user, offer.UserID)

	if userInfo.Error != nil {
		var resp = map[string]interface{}{"status": false, "message": "User Id not found in DB"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	createdOffer := db.Debug().Create(offer)
	var errMessage = createdOffer.Error

	if createdOffer.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdOffer)
}

func GetUserOffers(w http.ResponseWriter, r *http.Request) {
	var id = r.Context().Value("user")
	log.Printf("user : %v",id)
	var userOffers []models.Offers
	log.Printf("user info : %v",userOffers)
	offers := db.Debug().Where("user_id = ?", id).Find(&userOffers)

	json.NewEncoder(w).Encode(offers)
}
