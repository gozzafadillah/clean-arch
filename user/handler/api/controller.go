package userApi

import (
	"errors"
	"net/http"

	UserDomain "github.com/gozzafadillah/user/domain"
	"github.com/labstack/echo/v4"

	"github.com/gozzafadillah/user/handler/api/request"
	"github.com/gozzafadillah/user/handler/api/response"
)

type UserHandler struct {
	service UserDomain.Service
}

func NewUserHandler(service UserDomain.Service) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (uh UserHandler) Create(c echo.Context) error {
	req := request.RequestJSON{}
	if err := c.Bind(&req); err != nil {
		return errors.New("data gaada")
	}

	responseData, errResp := uh.service.InsertData(request.ToDomain(req))
	if errResp != nil {
		return errors.New("data error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    response.FromDomain(responseData),
		"message": "Insert success",
	})

}

func (uh UserHandler) Login(c echo.Context) error {
	req := request.RequestJSON{}

	if err := c.Bind(&req); err != nil {
		return errors.New("data gaada")
	}

	token, err := uh.service.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "anda tidak valid",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    token,
		"message": "JWT anda selesai",
	})
}
