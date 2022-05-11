package productApi

import (
	"fmt"
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
	fmt.Println("req :", req)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	responseData, err := ph.Service.CreateProduct(request.ToDomain(req))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Insert success",
		"rescode": http.StatusOK,
		"data":    response.FromDomain(responseData),
	})
}

func (ph *ProductHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	responseData, err := ph.Service.DestroyProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete success",
		"rescode": http.StatusOK,
		"data":    response.FromDomain(responseData),
	})
}

func (ph *ProductHandler) Update(c echo.Context) error {
	rec := request.RequestJSON{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&rec); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	response, err := ph.Service.EditProduct(id, request.ToDomain(rec))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "edit success",
		"rescode": http.StatusOK,
		"data":    response,
	})

}

func (ph *ProductHandler) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response, err := ph.Service.GetProductId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get product by id success",
		"rescode": http.StatusOK,
		"data":    response,
	})
}

func (ph *ProductHandler) GetAllProduct(c echo.Context) error {
	response, err := ph.Service.GetProducts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get all product success",
		"rescode": http.StatusOK,
		"data":    response,
	})
}
func (ph *ProductHandler) FilterPrice(c echo.Context) error {
	trigger := c.QueryParam("filter")
	// Min price
	if trigger == "min" {
		response, err := ph.Service.GetMinPrice()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "bad request",
				"rescode": http.StatusBadRequest,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "min price success",
			"rescode": http.StatusOK,
			"data":    response,
		})
		// Max price
	} else if trigger == "max" {
		response, err := ph.Service.GetMaxPrice()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "bad request",
				"rescode": http.StatusBadRequest,
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "max price success",
			"rescode": http.StatusOK,
			"data":    response,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "query parameter not match",
		"rescode": http.StatusBadRequest,
	})
}

func (ph *ProductHandler) CreateCategory(c echo.Context) error {
	req := request.RequestJSONCategory{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed in json request",
			"rescode": http.StatusBadRequest,
		})
	}

	response, err := ph.Service.CreateCategory(request.ToDomainCategory(req))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
			"rescode": http.StatusBadRequest,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "create category product success",
		"rescode": http.StatusOK,
		"data":    response,
	})
}

func (ph *ProductHandler) Category(c echo.Context) error {
	category := c.QueryParam("category")
	if category == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please insert query parameter",
		})
	}

	response, err := ph.Service.GetCategory(category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
			"rescode": http.StatusBadRequest,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "category " + category + " success",
		"rescode": http.StatusOK,
		"data":    response,
	})
}

func (ph *ProductHandler) GetCategoryById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	response, err := ph.Service.GetCategoryById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
			"rescode": echo.ErrBadRequest,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get category",
		"rescode": http.StatusOK,
		"data":    response,
	})
}
