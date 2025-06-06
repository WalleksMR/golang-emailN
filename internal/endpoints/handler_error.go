package endpoints

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/go-chi/render"
	"github.com/walleksmr/golang-emailn/internal/excptions"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpointFunc(w, r)

		if err != nil {
			if errors.Is(err, excptions.ErrInternal) {
				w.WriteHeader(500)
			} else {
				w.WriteHeader(400)
			}

			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		w.WriteHeader(status)
		v := reflect.ValueOf(obj)

		if obj != nil && !v.IsNil() {
			render.JSON(w, r, obj)
		}
	})
}
