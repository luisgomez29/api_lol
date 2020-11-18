package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/api_lol/controllers"
	"github.com/luisgomez29/api_lol/middlewares"
)

type CharacterRouter interface {
	CharacterRouters(e *echo.Group)
}

type characterRoutes struct {
	characterController controllers.CharacterController
}

func NewCharacterRoutes(characterController controllers.CharacterController) CharacterRouter {
	return &characterRoutes{characterController}
}

func (cr *characterRoutes) CharacterRouters(e *echo.Group) {
	e.GET("characters", cr.characterController.GetAll)
	e.GET("characters/:id", cr.characterController.FindById)
	e.POST("characters", cr.characterController.Create, middlewares.Authenticated())
	e.PUT("characters/:id", cr.characterController.Update, middlewares.Authenticated())
	e.DELETE("characters/:id", cr.characterController.Delete, middlewares.Authenticated())
}
