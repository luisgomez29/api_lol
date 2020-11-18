package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/luisgomez29/api_lol/config"
)

func Authenticated() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    config.SECRETKEY,
		SigningMethod: "HS512",
	})
}
