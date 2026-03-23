package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"user-mapping/domain/dto"
)

func GlobalExceptionHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				// log actual error
				log.Printf("[PANIC] %v", err)

				// always return 500
				writeError(w, http.StatusInternalServerError, "Internal server error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := dto.BaseResponseDto[any]{
		Result: dto.ResultResponseDto{
			Flag:        0,
			FlagMessage: message,
		},
	}

	json.NewEncoder(w).Encode(response)
}
