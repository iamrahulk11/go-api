package routes

import (
	"user-mapping/api"
	contract "user-mapping/domain/dto"
	helper "user-mapping/helper"
	"user-mapping/internal/container"

	"github.com/gin-gonic/gin"
)

// ParamMode defines how request parameters are passed
type ParamMode string

const (
	FormQuery ParamMode = "formquery"
	FormForm  ParamMode = "formform"
	FromBody  ParamMode = "frombody"
)

// Route defines a single API route
type Route struct {
	Path        string
	Method      string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
	Auth        bool
}

// CentralizedRoutes returns all routes for the app
func CentralizedRoutes(services *container.ServiceContainer, jwtHelper *helper.JWT) []Route {
	return []Route{
		{
			Path:   "/login",
			Method: "POST",
			Handler: helper.BindJsonRequestAndValidate(func(c *gin.Context, req contract.VerifyLoginRequest) {
				api.LoginHandler(services.LoginService).VerifyUser(c, req)
			}),
			Auth: false,
		},
		{
			Path:   "/user",
			Method: "GET",
			Handler: func(c *gin.Context) {
				api.UserHandler(services.UserService).User(c, c.Request)
			},
			Auth: true,
		},
		{
			Path:   "/profile",
			Method: "GET",
			Handler: helper.BindFromQueryRequestAndValidate(func(c *gin.Context, req contract.FetchUserProfileRequest) {
				api.UserHandler(services.UserService).FetchUserProfile(c, req)
			}),
			Auth: true,
		},
	}
}
