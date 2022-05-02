package main

import (
	"github.com/gozzafadillah/app/config"
	_middleware "github.com/gozzafadillah/app/middlewares"
	migrate "github.com/gozzafadillah/migrator"
	"github.com/gozzafadillah/routes"
	"github.com/labstack/echo/v4"

	userAPI "github.com/gozzafadillah/user/handler/api"
	userRepoMysql "github.com/gozzafadillah/user/repository/mysql"
	userService "github.com/gozzafadillah/user/service"
)

func main() {

	db := config.InitDB()
	migrate.AutoMigrate(db)

	configJWT := _middleware.ConfigJwt{
		SecretJWT: "2345",
	}

	e := echo.New()

	// Factory
	userRepo := userRepoMysql.NewUserRepository(db)
	userServ := userService.NewUserService(userRepo, &configJWT)
	userHandler := userAPI.NewUserHandler(userServ)

	//Route
	routesInit := routes.ControllerList{
		JWTMiddleware: configJWT.Init(),
		UserHandler:   userHandler,
	}
	routesInit.RouteRegister(e)

	e.Logger.Fatal(e.Start(":8080"))
}
