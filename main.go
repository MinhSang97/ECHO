package main

import (
	"app/dbutil"
	"app/handler"
	"app/helper"
	repoimpl "app/repository/repo_impl"
	"app/router"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"time"
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

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()

	e.Validator = structValidator

	// Lấy ngày hiện tại và định dạng thành chuỗi "2006-01-02"
	currentDate := time.Now().Format("2006-01-02")

	// Tạo đường dẫn thư mục và tên tệp tin log
	logDirectory := "log_middleware"
	logFileName := logDirectory + "/log_middleware_" + currentDate + ".txt"

	// Kiểm tra xem thư mục log có tồn tại không, nếu không tạo mới
	if _, err := os.Stat(logDirectory); os.IsNotExist(err) {
		err := os.Mkdir(logDirectory, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Mở hoặc tạo một tệp tin để lưu log middleware
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Sử dụng middleware.Logger() để ghi log và đặt Output là tệp tin đã mở
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
	}))

	userHandler := handler.UserHandler{
		UserRepo: repoimpl.NewUserRepo(sql),
	}

	repoHandler := handler.RepoHandler{
		GithubRepo: repoimpl.NewGithubRepo(sql),
	}

	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}

	api.SetupRouter()

	go scheduleUpdateTrending(5*time.Second, repoHandler)

	e.Logger.Fatal(e.Start(":3000"))
}

func scheduleUpdateTrending(timeSchedule time.Duration, handler handler.RepoHandler) {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Checking from github...")
				helper.CrawlRepo(handler.GithubRepo)
			}
		}
	}()
}
