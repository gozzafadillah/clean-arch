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

	productApi "github.com/gozzafadillah/product/handler/api"
	productRepoMysql "github.com/gozzafadillah/product/repository/mysql"
	productService "github.com/gozzafadillah/product/service"
)

func main() {

	db := config.InitDB()
	migrate.AutoMigrate(db)

	configJWT := _middleware.ConfigJwt{
		SecretJWT: "2345",
	}

	e := echo.New()

	// Factory
	// User
	userRepo := userRepoMysql.NewUserRepository(db)
	userServ := userService.NewUserService(userRepo, &configJWT)
	userHandler := userAPI.NewUserHandler(userServ)
	// product
	productRepo := productRepoMysql.NewProductRepository(db)
	productServ := productService.NewProductService(productRepo)
	productHandler := productApi.NewProductHandler(productServ)

	//Route
	routesInit := routes.ControllerList{
		JWTMiddleware:  configJWT.Init(),
		UserHandler:    userHandler,
		ProductHandler: productHandler,
	}
	routesInit.RouteRegister(e)

	e.Logger.Fatal(e.Start(":8080"))
}
