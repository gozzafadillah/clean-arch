package main

import (
	"github.com/gozzafadillah/app/config"
	"github.com/gozzafadillah/app/middlewares"
	migrate "github.com/gozzafadillah/migrator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/gozzafadillah/user"
)

const jwtToken = "123456"

func main() {

	db := config.InitDB()
	migrate.AutoMigrate(db)

	user := user.NewUserFactory(db)
	//Route
	e := echo.New()

	middlewares.LogMiddleware(e)

	auth := e.Group("")
	auth.Use(middleware.JWT(middlewares.ConfigJwt{
		SecretJWT: jwtToken,
	}))
	auth.POST("/users", user.Create)

	e.POST("/login", user.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
