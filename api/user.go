package api

import (
	"net/http"
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
	result := h.services.UserService()
	status := 200
	if result.Result.Flag == 0 {
		status = 400
	}
	c.JSON(status, result)
}

func (h *UserHandlerStruct) FetchUserProfile(c *gin.Context, req profileRequest.FetchUserProfileRequestDto) {
	// Call service
	result := h.services.FetchUserProfileDetails(req)
	status := 200
	if result.Result.Flag == 0 {
		status = 400
	}
	c.JSON(status, result)
}
