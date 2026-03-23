package helper

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

type HandlerWithDTO[T any] func(w http.ResponseWriter, dto T)

func WrapHandlerWithDTO[T any](handler HandlerWithDTO[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto T
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		handler(w, dto)
	}
}

func WrapHandlerWithQuery[T any](handler func(http.ResponseWriter, T)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto T
		decoder := schema.NewDecoder()
		if err := decoder.Decode(&dto, r.URL.Query()); err != nil {
			http.Error(w, "invalid query params", http.StatusBadRequest)
			return
		}
		handler(w, dto)
	}
}
