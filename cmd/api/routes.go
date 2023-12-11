package main

import (
	"net/http"

	contract "authentication/contract"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "authentication/cmd/api/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetRoutes(e *echo.Echo, userService contract.UserService) http.Handler {
	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://*", "https://*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentType, echo.HeaderXCSRFToken},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Heartbeat route
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// Authentication route
	e.POST("/authenticate", func(c echo.Context) error {
		return HandlerAuthenticate(c, userService)
	})

	return e
}
