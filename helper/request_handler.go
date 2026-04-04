package helper

import (
	"fmt"
	base_response "user-mapping/domain/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func BindJsonRequestAndValidate[T any](handler func(*gin.Context, T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto T

		// Bind JSON
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(400, base_response.BaseResponseDto[any]{
				Result: base_response.ResultResponseDto{
					Flag:        0,
					FlagMessage: err.Error(),
				},
			})
			return
		}

		// Validate required fields
		if err := validate.Struct(dto); err != nil {
			c.JSON(400, base_response.BaseResponseDto[any]{
				Result: base_response.ResultResponseDto{
					Flag:        0,
					FlagMessage: ValidationErrorToMessage(err),
				},
			})
			return
		}

		// Call actual handler with bound and validated DTO
		handler(c, dto)
	}
}

func BindFromQueryRequestAndValidate[T any](handler func(*gin.Context, T)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto T
		if err := c.ShouldBindQuery(&dto); err != nil {
			c.JSON(400, base_response.BaseResponseDto[any]{
				Result: base_response.ResultResponseDto{
					Flag:        0,
					FlagMessage: ValidationErrorToMessage(err),
				},
			})
			return
		}

		if err := validate.Struct(dto); err != nil {
			c.JSON(400, base_response.BaseResponseDto[any]{
				Result: base_response.ResultResponseDto{
					Flag:        0,
					FlagMessage: ValidationErrorToMessage(err),
				},
			})
			return
		}

		handler(c, dto)
	}
}

// Returns a user-friendly error message
func ValidationErrorToMessage(err error) string {
	if err == nil {
		return ""
	}

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			key := fmt.Sprintf("%s.%s", e.Field(), e.Tag())
			if msg, exists := customMessages[key]; exists {
				return msg
			}
			// default fallback
			return fmt.Sprintf("Field %s failed validation on %s", e.Field(), e.Tag())
		}
	}
	return err.Error()
}
