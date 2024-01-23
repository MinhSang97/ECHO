package main

import (
	"app/dbutil"
	"app/handler"
	repoimpl "app/repository/repo_impl"
	"app/router"
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
	userHandler := handler.UserHandler{
		UserRepo: repoimpl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}
