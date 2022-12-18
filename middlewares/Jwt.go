package middlewares

import (
	"net/http"

	"github.com/Surrendra/auth-go/config"
	"github.com/Surrendra/auth-go/helper"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{
					"message": "Anda tidak memiliki ijin",
				}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			}
		}

		tokenString := c.Value
		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				response := map[string]string{
					"message": "Token tidak valid",
				}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				response := map[string]string{
					"message": "Token yang anda masukan sudah expired",
				}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{
					"message": "Internal Server Error",
				}
				helper.ResponseJson(w, http.StatusInternalServerError, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{
				"message": "Token tidak valid",
			}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
