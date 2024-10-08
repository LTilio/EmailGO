package endpoints

import (
	internalerror "EmailGO/internal/internalError"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type Endpoint func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(handler Endpoint) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := handler(w, r)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			if errors.Is(err, internalerror.ErrInternal) {
				render.Status(r, 500)
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				render.Status(r, 404)
			} else {
				render.Status(r, 400)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
