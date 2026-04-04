package api

import (
	"net/http"
	"user-mapping/domain/dto"
	profileRequest "user-mapping/domain/dto/requests/user"
	"user-mapping/domain/services"

	"github.com/gin-gonic/gin"
)

type UserHandlerStruct struct {
	services *services.UserServiceStruct
}

func UserHandler(services *services.UserServiceStruct) *UserHandlerStruct {
	return &UserHandlerStruct{services: services}
}

// VerifyUser now accepts the parsed request DTO
func (h *UserHandlerStruct) User(c *gin.Context, r *http.Request) {
	result, err := h.services.UserService()
	if err != nil {
		c.JSON(400, dto.BaseResponseDto[any]{
			Result: dto.ResultResponseDto{
				Flag:        0,
				FlagMessage: err.Error(),
			},
		})
		return
	}

	// Success
	c.JSON(200, dto.BaseResponseDto[any]{
		Result: dto.ResultResponseDto{
			Flag:        1,
			FlagMessage: "Success",
		},
		Data: result,
	})
}

func (h *UserHandlerStruct) FetchUserProfile(c *gin.Context, req profileRequest.FetchUserProfileRequestDto) {
	// Call service
	result, err := h.services.FetchUserProfileDetails(req)
	if err != nil {
		c.JSON(400, dto.BaseResponseDto[any]{
			Result: dto.ResultResponseDto{
				Flag:        0,
				FlagMessage: err.Error(),
			},
		})
		return
	}

	// Success
	c.JSON(200, dto.BaseResponseDto[any]{
		Result: dto.ResultResponseDto{
			Flag:        1,
			FlagMessage: "Success",
		},
		Data: result,
	})
}
