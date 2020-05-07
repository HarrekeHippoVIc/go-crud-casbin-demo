package middlewares

import (
	"net/http"

	"github.com/HarrekeHippoVic/go-crud-casbin-demo/api/auth"
	"github.com/HarrekeHippoVic/go-crud-casbin-demo/api/responses"
	"github.com/pkg/errors"
)

// SetMiddlewareJSON func
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentificaton func
func SetMiddlewareAuthentificaton(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
