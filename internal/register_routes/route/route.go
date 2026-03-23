package routes

import (
	"net/http"
	"user-mapping/api"
	"user-mapping/helper"
	"user-mapping/internal/container"
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
	Handler     http.HandlerFunc
	Middlewares []func(http.Handler) http.Handler
	ParamMode   ParamMode
	Auth        bool
	DTO         any
}

// CentralizedRoutes returns all routes for the app
func CentralizedRoutes(services *container.ServiceContainer, jwtHelper *helper.JWT) []Route {

	return []Route{
		{
			Path:        "/login",
			Method:      "POST",
			Handler:     helper.WrapHandlerWithDTO(api.LoginHandler(services.LoginService).VerifyUser),
			Middlewares: []func(http.Handler) http.Handler{},
			Auth:        false,
		},
		{
			Path:        "/user",
			Method:      "GET",
			Handler:     api.UserHandler(services.UserService).User,
			Middlewares: []func(http.Handler) http.Handler{},
			Auth:        true,
		},
		{
			Path:   "/profile",
			Method: "GET",
			Handler: helper.WrapQueryHandler(
				api.UserHandler(services.UserService).FetchUserProfile,
			),
			Middlewares: []func(http.Handler) http.Handler{},
			Auth:        true,
		},
	}
}
