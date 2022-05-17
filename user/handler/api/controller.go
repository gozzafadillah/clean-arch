package userApi

import (
	"errors"
	"net/http"

	"github.com/gozzafadillah/app/middlewares"
	userDomain "github.com/gozzafadillah/user/domain"
	"github.com/labstack/echo/v4"

	"github.com/gozzafadillah/user/handler/api/request"
	"github.com/gozzafadillah/user/handler/api/response"
)

type UserHandler struct {
	service userDomain.Service
}

func NewUserHandler(service userDomain.Service) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (uh *UserHandler) Create(c echo.Context) error {
	req := request.RequestJSON{}
	if err := c.Bind(&req); err != nil {
		return errors.New("data invalid")
	}

	responseData, errResp := uh.service.InsertData(request.ToDomain(req))
	if errResp != nil {
		return errors.New("data invalid")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    response.FromDomain(responseData),
		"message": "insert success",
	})

}

func (uh *UserHandler) Login(c echo.Context) error {
	req := request.RequestJSON{}

	if err := c.Bind(&req); err != nil {
		return errors.New("data invalid")
	}

	token, err := uh.service.Login(req.Username, req.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "invalid login",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success login",
		"response": http.StatusOK,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

func (uh *UserHandler) BanUser(c echo.Context) error {
	username := c.Param("username")

	// Check user jwt
	checkClaim := middlewares.GetUser(c)

	// mencegah ban sesama admin
	userInput, err := uh.service.GetId(checkClaim.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "user input unknown",
		})
	}
	user, err := uh.service.GetUsername(username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "user from request unknown",
		})
	}
	if user.Role == userInput.Role {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "dont ban same admin",
		})
	}

	// Ban User
	res, err := uh.service.BanUser(username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "username wrong",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success ban user",
		"rescode": http.StatusOK,
		"data":    response.FromDomain(res),
	})
}

func (uh *UserHandler) UserRole(id int) (string, bool) {
	var role string
	var status bool
	user, err := uh.service.GetId(id)
	if err == nil {
		role = user.Role
		status = user.Status
	}
	return role, status
}
