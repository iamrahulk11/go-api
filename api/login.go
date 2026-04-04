package api

import (
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
	result := h.services.VerifyUserService(loginReq)
	status := 200
	if result.Result.Flag == 0 {
		status = 400
	}
	c.JSON(status, result)
}
