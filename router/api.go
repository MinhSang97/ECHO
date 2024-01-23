package router

import (
	"app/handler"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
	//RepoHandler handler.RepoHandler
}

func (api *API) SetupRouter() {
	// user
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
}
