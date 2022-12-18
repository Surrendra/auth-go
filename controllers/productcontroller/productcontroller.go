package productcontroller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Surrendra/auth-go/helper"
	"github.com/Surrendra/auth-go/models"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	if err := models.DB.Find(&products).Error; err != nil {
		fmt.Println(err)
	}
	response := map[string]interface{}{
		"data": products,
	}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		response := map[string]string{
			"message": "id tidak ditemukan pada database",
		}
		helper.ResponseJson(w, http.StatusBadRequest, response)
		return
	}

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		response := map[string]string{
			"message": "Data yang anda cari tidak ditemukan",
		}
		helper.ResponseJson(w, http.StatusNotFound, response)
		return
	}
	jumlah := 12 * 30
	response := map[string]interface{}{
		"message":         "Success",
		"data":            product,
		"additional_data": jumlah,
	}
	helper.ResponseJson(w, http.StatusOK, response)
}
