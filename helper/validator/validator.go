package validator

import (
	"errors"
	"fmt"

	"github.com/gozzafadillah/app/middlewares"
	userApi "github.com/gozzafadillah/user/handler/api"
	"github.com/labstack/echo/v4"
)

func RoleValidation(role string, userControler userApi.UserHandler) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewares.GetUser(c)
			fmt.Println("claim : ", claims)
			userRole := userControler.UserRole(claims.ID)
			fmt.Println("userRole : ", userRole)

			if userRole == role {
				return hf(c)
			} else {
				return errors.New("role forbiden")
			}
		}
	}
}
