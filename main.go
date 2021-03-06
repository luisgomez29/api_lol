package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/luisgomez29/api_lol/auto"
	"github.com/luisgomez29/api_lol/config"
	"github.com/luisgomez29/api_lol/controllers"
	"github.com/luisgomez29/api_lol/database"
	"github.com/luisgomez29/api_lol/repositories"
	"github.com/luisgomez29/api_lol/routes"
)

var resetTables = flag.Bool("rt", false, "Reset tables")

func main() {
	flag.Parse()
	config.Load()
	db := database.Connect()
	if db != nil {
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
	}

	if *resetTables {
		auto.Load(db)
	}

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)
	userRoutes := routes.NewUserRoutes(userController)

	characterRepository := repositories.NewCharacterRepository(db)
	characterController := controllers.NewCharacterController(characterRepository)
	characterRouter := routes.NewCharacterRoutes(characterController)

	loginController := controllers.NewAuthController(db)
	loginRouter := routes.NewAuthRouter(loginController)

	e := echo.New()
	apiV1 := e.Group("/api/v1/")

	//apiV1.Use(middlewares.Authenticated())

	routes.InitRoutes(apiV1, userRoutes, characterRouter, loginRouter)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.PORT)))
}
