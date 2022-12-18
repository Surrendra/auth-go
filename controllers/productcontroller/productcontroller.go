package productcontroller

import (
	"net/http"

	"github.com/Surrendra/auth-go/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":    1,
			"name":  "Kemeja",
			"stock": 1000,
		},
		{
			"id":    2,
			"name":  "Celana",
			"stock": 8000,
		},
		{
			"id":    3,
			"name":  "CD Gratis Sepatu",
			"stock": 9000,
		},
	}

	helper.ResponseJson(w, http.StatusOK, data)
}
