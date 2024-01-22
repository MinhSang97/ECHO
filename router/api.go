package router

import (
	"app/handler"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func SetupRouter() {
	API.Echo.GET("/user/sign-in", api.UserHandler.HandleSignIn)
	API.Echo.GET("/user/sign-up", api.UserHandler.HandleSignUp)
}
