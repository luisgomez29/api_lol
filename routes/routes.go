package routes

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Group, ur UserRouter, pr CharacterRouter, lr AuthRouter) {
	ur.UserRouters(e)
	pr.CharacterRouters(e)
	lr.AuthRoutes(e)
}
