package main

import (
	"github.com/gozzafadillah/app/config"
	migrate "github.com/gozzafadillah/migrator"
	"github.com/labstack/echo/v4"

	"github.com/gozzafadillah/user"
)

func main() {

	db := config.InitDB()
	migrate.AutoMigrate(db)

	user := user.NewUserFactory(db)
	//Route
	e := echo.New()

	e.POST("/users", user.Create)
	e.POST("/login", user.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
