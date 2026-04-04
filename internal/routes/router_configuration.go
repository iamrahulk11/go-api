package routes

import (
	"net/http"
	"user-mapping/helper"
	"user-mapping/internal/container"
	middlewares "user-mapping/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RouteRegistry stores all routes
type RouteRegistry struct {
	Routes []Route
}

// RegisterAppRoutes initializes router and registers all centralized routes
func RegisterAppRoutes(services *container.ServiceContainer, jwtHelper *helper.JWT) *gin.Engine {
	router := gin.New()

	// global middlewares
	router.Use(LoggingMiddleware(), middlewares.GlobalExceptionHandler())

	registry := &RouteRegistry{}

	allRoutes := CentralizedRoutes(services, jwtHelper)

	registry.Routes = append(registry.Routes, allRoutes...)

	registry.RegisterAll(router, jwtHelper)

	return router
}

func (r *RouteRegistry) RegisterAll(router *gin.Engine, jwtHelper *helper.JWT) {

	for _, route := range r.Routes {

		handlers := []gin.HandlerFunc{}

		// default middleware
		handlers = append(handlers, LoggingMiddleware(), middlewares.GlobalExceptionHandler())

		// route-specific middleware
		handlers = append(handlers, route.Middlewares...)

		// auth middleware
		if route.Auth {
			handlers = append(handlers, middlewares.JWTMiddleware(jwtHelper))
		}

		// final handler
		handlers = append(handlers, route.Handler)

		// register based on method
		switch route.Method {
		case "GET":
			router.GET(route.Path, handlers...)
		case "POST":
			router.POST(route.Path, handlers...)
		case "PUT":
			router.PUT(route.Path, handlers...)
		case "DELETE":
			router.DELETE(route.Path, handlers...)
		default:
			router.Any(route.Path, handlers...)
		}
	}
}

// LoggingMiddleware logs basic info about requests
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Request:", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

// custom response writer
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
