package routes

import (
	"github.com/gozzafadillah/app/middlewares"
	userApi "github.com/gozzafadillah/user/handler/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware middleware.JWTConfig
	UserHandler   userApi.UserHandler
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	middlewares.LogMiddleware(e)

	e.POST("/login", cl.UserHandler.Login)
	users := e.Group("users")
	users.Use(middleware.JWTWithConfig(cl.JWTMiddleware))
	users.POST("/register", cl.UserHandler.Create)

}
