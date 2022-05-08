package productApi

import (
	"net/http"
	"strconv"

	productDomain "github.com/gozzafadillah/product/domain"
	"github.com/gozzafadillah/product/handler/api/request"
	"github.com/gozzafadillah/product/handler/api/response"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Service productDomain.Service
}

func NewProductHandler(service productDomain.Service) ProductHandler {
	return ProductHandler{
		Service: service,
	}
}

func (ph *ProductHandler) Create(c echo.Context) error {
	req := request.RequestJSON{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad request",
		})
	}

	responseData, err := ph.Service.CreateProduct(request.ToDomain(req))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    response.FromDomain(responseData),
		"message": "Insert success",
	})
}

func (ph *ProductHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	responseData, err := ph.Service.DestroyProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    response.FromDomain(responseData),
		"message": "delete success",
	})
}

func (ph *ProductHandler) Update(c echo.Context) error {
	rec := request.RequestJSON{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&rec); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}

	response, err := ph.Service.EditProduct(id, request.ToDomain(rec))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "edit success",
		"data":    response,
	})

}

func (ph *ProductHandler) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := ph.Service.GetProductId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "edit success",
		"data":    response,
	})
}

func (ph *ProductHandler) GetAllProduct(c echo.Context) error {
	response, err := ph.Service.GetProducts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "edit success",
		"data":    response,
	})
}
