package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Surrendra/auth-go/helper"
	"github.com/Surrendra/auth-go/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	var request models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		log.Fatal("Gagal decode json")
	}
	defer r.Body.Close()
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(hashPassword)
	if err := models.DB.Create(&request).Error; err != nil {
		log.Fatal("Gagal menyimpan data ke database")
		helper.ResponseJson(w, http.StatusBadRequest, err.Error())
		return
	}
	response := map[string]interface{}{
		"message": "Success",
		"data":    nil,
	}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
