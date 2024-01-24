package middleware

import (
	"app/model"
	sercurity2 "app/sercurity"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		SigningKey: []byte(sercurity2.SECRET_KEY),
		Claims:     &model.JwtCustomClaims{},
	}

	return middleware.JWTWithConfig(config)
}
