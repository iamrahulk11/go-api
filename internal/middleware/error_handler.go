package middlewares

import (
	"net/http"
	"user-mapping/domain/dto"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0].Error()

			response := dto.BaseResponseDto[any]{
				Result: dto.ResultResponseDto{
					Flag:        0,
					FlagMessage: err,
				},
			}

			c.JSON(http.StatusBadRequest, response)
			c.Abort()
		}
	}
}
