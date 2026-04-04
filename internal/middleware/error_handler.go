package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"user-mapping/domain/dto"

	"github.com/gin-gonic/gin"
)

func GlobalExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				// log actual error
				log.Printf("[PANIC] %v", err)

				// always return 500
				writeError(c.Writer, http.StatusInternalServerError, "Internal server error")
			}
		}()

		c.Next()
	}
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
