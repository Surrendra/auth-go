package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Surrendra/auth-go/config"
	"github.com/Surrendra/auth-go/helper"
	"github.com/Surrendra/auth-go/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var request models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		log.Fatal("Gagal decode json")
	}
	defer r.Body.Close()

	// check username
	var user models.User
	if err := models.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{
				"message": "Username tidak ditemukan !",
			}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{
				"message": err.Error(),
			}
			helper.ResponseJson(w, http.StatusInternalServerError, response)
			return
		}
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		response := map[string]string{
			"message": "Password yang anda masukan salah !",
		}
		helper.ResponseJson(w, http.StatusUnauthorized, response)
		return
	}

	expTime := time.Now().Add(time.Minute * 30)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "auth-go",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgorithm.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]interface{}{
			"message": err.Error(),
			"token":   config.JWT_KEY,
		}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// update api token in user table
	if models.DB.Model(&user).Where("id", user.Id).Update("api_token", token).RowsAffected == 0 {
		response := map[string]interface{}{
			"message": "Gagal melakukan update token ke table users",
			"token":   config.JWT_KEY,
		}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	// set token to cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})
	response := map[string]interface{}{
		"message": "Anda berhasil login",
		"data":    user,
		"token":   token,
	}
	helper.ResponseJson(w, http.StatusOK, response)

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
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
	response := map[string]interface{}{
		"message": "Logout berhasil",
		"data":    nil,
		"token":   nil,
	}
	helper.ResponseJson(w, http.StatusOK, response)
}
