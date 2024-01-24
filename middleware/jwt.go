package middleware

import (
	"app/model"
	"app/sercurity"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: sercurity.SECRET_KEY,
	}
	return middleware.JWTWithConfig(config)
}
