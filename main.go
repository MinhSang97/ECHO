package main

import (
	"app/dbutil"
	"app/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	sql := &dbutil.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "admin",
		PassWord: "123456",
		DbName:   "golang",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}
