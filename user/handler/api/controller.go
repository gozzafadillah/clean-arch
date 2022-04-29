package userApi

import (
	"errors"

	UserDomain "github.com/gozzafadillah/user/domain"
	"github.com/labstack/echo/v4"

	"github.com/gozzafadillah/user/handler/api/request"
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

	_, errResp := uh.service.InsertData(request.ToDomain(req))

	if errResp != nil {
		return errors.New("data error")
	}

	return nil

}
