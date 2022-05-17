package validator

import (
	"fmt"
	"net/http"

	"github.com/gozzafadillah/app/middlewares"
	userApi "github.com/gozzafadillah/user/handler/api"
	"github.com/labstack/echo/v4"
)

func RoleValidation(role string, userControler userApi.UserHandler) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewares.GetUser(c)
			fmt.Println("claim : ", claims)
			userRole, status := userControler.UserRole(claims.ID)
			fmt.Println("userRole : ", userRole)

			if userRole == role && status == true {
				return hf(c)
			} else {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "account not active, please contact admin",
				})
			}
		}
	}
}
