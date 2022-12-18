package productcontroller

import (
	"encoding/json"
	"fmt"
	"log"
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

func Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		log.Fatal("gagal melakukan decode request data")
	}
	defer r.Body.Close()
	if err := models.DB.Create(&product).Error; err != nil {
		log.Fatal("gagal mennyimpan data")
		response := map[string]string{
			"message": "Terjadi kesalahan ketika menyimpan data",
		}
		helper.ResponseJson(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]interface{}{
		"message": "Berhasil menyimpan data",
		"data":    product,
	}
	helper.ResponseJson(w, http.StatusOK, response)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		log.Fatal("gagal melakukan decode request data")
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Fatal("terjadi kesalahn ketika mengambil ID data")
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		log.Fatal("Terjadi kesalahan saat update data")
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Fatal("terjadi kesalahn ketika mengambil ID data")
	}
	var FindProduct models.Product
	if err := models.DB.First(&FindProduct, id).Error; err != nil {
		response := map[string]string{
			"message": "Product tidak ditemukan atau sudah terhapus !",
		}
		helper.ResponseJson(w, http.StatusNotFound, response)
		return
	}

	var DeleteProduct models.Product
	if models.DB.Delete(DeleteProduct, id).RowsAffected == 0 {
		response := map[string]string{
			"message": "Product tidak ditemukan atau sudah terhapus !",
		}
		helper.ResponseJson(w, http.StatusNotFound, response)
		return
	}
	response := map[string]string{
		"message": "Product " + FindProduct.Name + " Dihapus !",
	}
	helper.ResponseJson(w, http.StatusOK, response)
}
