package helper

import (
	"encoding/json"
	"net/http"
	"user-mapping/domain/dto"
	"user-mapping/domain/validator"

	"github.com/gorilla/schema"
)

type HandlerWithDTO[T any] func(w http.ResponseWriter, dto T)

func WrapHandlerWithDTO[T any](handler func(http.ResponseWriter, T)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dtoReq T
		w.Header().Set("Content-Type", "application/json")

		// Decode
		if err := json.NewDecoder(r.Body).Decode(&dtoReq); err != nil {
			writeError(w, "invalid request body")
			return
		}

		// Validate
		if err := validator.Validate.Struct(dtoReq); err != nil {
			validationErrors := validator.FormatValidationErrors(err)

			msg := ""
			for _, v := range validationErrors {
				msg += v + "; "
			}

			writeError(w, msg)
			return
		}

		handler(w, dtoReq)
	}
}

func WrapQueryHandler[T any](handler func(http.ResponseWriter, *http.Request, T)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var dto T

		decoder := schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)

		// Bind
		if err := decoder.Decode(&dto, r.URL.Query()); err != nil {
			//writeError(w, "invalid query parameters")
			//return
		}

		// Validate
		if err := validator.Validate.Struct(dto); err != nil {
			validationErrors := validator.FormatValidationErrors(err)

			msg := ""
			for _, v := range validationErrors {
				msg += v + "; "
			}

			writeError(w, msg)
			return
		}

		// ✅ Pass validated DTO to handler
		handler(w, r, dto)
	}
}

func writeError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)

	_ = json.NewEncoder(w).Encode(dto.BaseResponseDto[any]{
		Result: dto.ResultResponseDto{
			Flag:        0,
			FlagMessage: message,
		},
	})
}
