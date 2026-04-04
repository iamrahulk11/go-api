package api

import (
	"user-mapping/domain/dto"
	request "user-mapping/domain/dto/requests/login"
	"user-mapping/domain/services"

	"github.com/gin-gonic/gin"
)

type LoginHandlerStruct struct {
	services *services.LoginServiceStruct
}

func LoginHandler(services *services.LoginServiceStruct) *LoginHandlerStruct {
	return &LoginHandlerStruct{services: services}
}

// VerifyUser now accepts the parsed request DTO
func (h *LoginHandlerStruct) VerifyUser(c *gin.Context, loginReq request.VerifyLoginRequestDto) {
	// Call service
	result, err := h.services.VerifyUserService(loginReq)
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
