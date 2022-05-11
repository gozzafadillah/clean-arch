package routes

import (
	"github.com/gozzafadillah/app/middlewares"
	"github.com/gozzafadillah/helper/validator"
	productApi "github.com/gozzafadillah/product/handler/api"
	userApi "github.com/gozzafadillah/user/handler/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserHandler    userApi.UserHandler
	ProductHandler productApi.ProductHandler
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	middlewares.LogMiddleware(e)

	// user
	e.POST("/login", cl.UserHandler.Login)
	e.POST("/register", cl.UserHandler.Create)

	// product public
	product := e.Group("product")
	product.GET("/all", cl.ProductHandler.GetAllProduct)
	product.GET("/all", cl.ProductHandler.FilterPrice)
	product.GET("/", cl.ProductHandler.Category)
	product.GET("/:id", cl.ProductHandler.GetProduct)

	// admin only / private
	authProduct := e.Group("product")
	authProduct.Use(middleware.JWTWithConfig(cl.JWTMiddleware), validator.RoleValidation("admin", cl.UserHandler))
	authProduct.POST("/create", cl.ProductHandler.Create)
	authProduct.DELETE("/delete/:id", cl.ProductHandler.Delete)
	authProduct.PUT("/update/:id", cl.ProductHandler.Update)

	authCategory := e.Group("category")
	authCategory.Use(middleware.JWTWithConfig(cl.JWTMiddleware), validator.RoleValidation("admin", cl.UserHandler))
	authCategory.POST("/create", cl.ProductHandler.CreateCategory)
	authCategory.GET("/:id", cl.ProductHandler.GetCategoryById)
}
