package middleware

import (
	"app/model"
	"app/model/req"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ISAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//handle logic
			req := req.ReqSignIn{}
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			if req.Email != "minhsangnguyen463@gmail.com" {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "Bạn không quyền gọi API này",
					Data:       nil,
				})
			}
			return next(c)
		}
	}
}
