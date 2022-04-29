package main

import (
	"github.com/gozzafadillah/app/config"
	migrate "github.com/gozzafadillah/migrator"
	"github.com/labstack/echo/v4"

	"github.com/gozzafadillah/user"
)

func main() {

	db := config.InitDB()
	migrate.AutoMigrate()

	user := user.NewUserFactory(db)
	//Route
	e := echo.New()

	e.POST("/user", user.Create)

	e.Start(":8080")
}
