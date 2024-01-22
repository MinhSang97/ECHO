package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Sang",
		"email": "minhsangnguyen463@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		Email    string
		FullName string
		Age      int
	}
	user := User{
		Email:    "minhsangnguyen463@gmail.com",
		FullName: "Nguyễn Minh Sang",
		Age:      18,
	}
	return c.JSON(http.StatusOK, user)
}
