package handler

import (
	"app/model"
	"app/security"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req2.ReqSignUp{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	hash := security.HashAndSalt([]byte(req.PassWord))
	role := model.MEMBER.String()

	userID, err := uuid.NewUUID()

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Sang",
		"email": "minhsangnguyen463@gmail.com",
	})
}
