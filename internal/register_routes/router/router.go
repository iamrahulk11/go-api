package router

import (
	"net/http"
	"user-mapping/helper"
	"user-mapping/internal/container"
	middlewares "user-mapping/internal/middleware"
	routes "user-mapping/internal/register_routes/route"

	"github.com/go-chi/chi/v5"
)

// RouteRegistry stores all routes
type RouteRegistry struct {
	Routes []routes.Route
}

// RegisterAppRoutes initializes router and registers all centralized routes
func RegisterAppRoutes(services *container.ServiceContainer, jwtHelper *helper.JWT) *chi.Mux {
	router := chi.NewRouter()
	router.Use(LoggingMiddleware)
	router.Use(middlewares.GlobalExceptionHandler)

	registry := &RouteRegistry{}

	// get all routes from centralized route package
	allRoutes := routes.CentralizedRoutes(services, jwtHelper)

	// add all routes to registry
	for _, r := range allRoutes {
		registry.Routes = append(registry.Routes, r)
	}

	// register all routes with middlewares, auth, etc.
	registry.RegisterAll(router, jwtHelper)

	return router
}

// RegisterAll registers all routes into the Chi router
func (r *RouteRegistry) RegisterAll(router *chi.Mux, jwtHelper *helper.JWT) {
	// default middleware stack (logging, recovery)
	defaultMiddlewares := []func(http.Handler) http.Handler{
		LoggingMiddleware,
		middlewares.GlobalExceptionHandler,
	}

	for _, route := range r.Routes {
		var h http.Handler = route.Handler

		// apply default middlewares
		for _, mw := range defaultMiddlewares {
			h = mw(h)
		}

		// apply route-specific middlewares
		for _, mw := range route.Middlewares {
			h = mw(h)
		}

		// auto-apply auth middleware if Auth is true
		if route.Auth {
			jwtMiddleware := &middlewares.JWTMiddleware{JWT: jwtHelper}
			h = jwtMiddleware.Handle(h)
		}

		// register route based on method
		router.Method(route.Method, route.Path, h)
	}
}

// LoggingMiddleware logs basic info about requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("Request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
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
