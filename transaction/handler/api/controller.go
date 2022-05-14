package transactionApi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gozzafadillah/app/middlewares"
	productDomain "github.com/gozzafadillah/product/domain"
	transactionDomain "github.com/gozzafadillah/transaction/domain"
	"github.com/gozzafadillah/transaction/handler/api/request"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	ServiceTransaction transactionDomain.Service
	ServiceProduct     productDomain.Service
}

func NewTransactionHandler(transaction transactionDomain.Service, product productDomain.Service) TransactionHandler {
	return TransactionHandler{
		ServiceTransaction: transaction,
		ServiceProduct:     product,
	}
}

func (th *TransactionHandler) CreateData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// request body from user
	req := request.RequestJSONCheckout{}

	product, err := th.ServiceProduct.CheckoutProductId(id)
	fmt.Println("product :", product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
			"rescode": http.StatusBadRequest,
		})
	}

	// get user id
	claims := middlewares.GetUser(c)
	// buat checkout
	code := th.ServiceTransaction.GetCode()
	fmt.Println("code : ", code)
	// Memasukan bind body ke memory request
	if errReq := c.Bind(&req); errReq != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	respCheckout, err := th.ServiceTransaction.CreateCheckout(code, request.ToDomainCheckout(req), product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
			"rescode": http.StatusBadRequest,
		})
	}
	fmt.Println("checkout : ", respCheckout)

	// create transaction
	transaction, err := th.ServiceTransaction.CreateTransaction(claims.ID, code, respCheckout)
	fmt.Println("transaction : ", transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
			"rescode": http.StatusBadRequest,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Checkout success",
		"rescode": http.StatusOK,
		"data":    transaction,
	})
}
func (th *TransactionHandler) CreateOngkir(c echo.Context) error {
	id, err := th.ServiceTransaction.CheckCity("Bandung")
	fmt.Println("id :", id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Checkout success",
		"rescode": http.StatusOK,
		"data":    id,
	})
}
