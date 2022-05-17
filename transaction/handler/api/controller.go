package transactionApi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gozzafadillah/app/middlewares"
	productDomain "github.com/gozzafadillah/product/domain"
	transactionDomain "github.com/gozzafadillah/transaction/domain"
	"github.com/gozzafadillah/transaction/handler/api/request"
	"github.com/gozzafadillah/transaction/handler/api/response"
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
	// input
	id, _ := strconv.Atoi(c.Param("id"))

	// request body from user
	req := request.RequestJSONCheckout{}

	// Get product
	product, err := th.ServiceProduct.CheckoutProductId(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
			"rescode": http.StatusBadRequest,
		})
	}

	// cek quantity apakah melebihi kapasitas

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
	// Membuat Checkout
	respCheckout, err := th.ServiceTransaction.CreateCheckout(code, request.ToDomainCheckout(req), product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error, checkout fail",
			"rescode": http.StatusBadRequest,
		})
	}
	// origin
	originId, err := th.ServiceTransaction.CheckCity(product.Origin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "origin city not found",
			"rescode": http.StatusBadRequest,
		})
	}

	//Destination
	destinationId, err := th.ServiceTransaction.CheckCity(respCheckout.Destination)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "destination city not found",
			"rescode": http.StatusBadRequest,
		})
	}
	// Ongkir
	ongkirPrice, etd, err := th.ServiceTransaction.Ongkir(originId, destinationId, int(respCheckout.Weight), respCheckout.Courier, respCheckout.Package)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "courier and package not found",
			"rescode": http.StatusBadRequest,
		})
	}

	// create transaction
	transaction, err := th.ServiceTransaction.CreateTransaction(claims.ID, code, ongkirPrice, etd, respCheckout)
	fmt.Println("transaction : ", transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "transaction fail",
			"rescode": http.StatusBadRequest,
		})
	}

	// mengurangi qty di product

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "transaction success",
		"rescode": http.StatusOK,
		"data":    response.FromDomainCheckout(transaction),
	})
}

// func (th *TransactionHandler) CreateOngkir(c echo.Context) error {
// 	id, err := th.ServiceTransaction.CheckCity("Jakarta Tenggara")
// 	fmt.Println("id :", id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": "City not found",
// 			"rescode": http.StatusBadRequest,
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Checkout success",
// 		"rescode": http.StatusOK,
// 		"data":    id,
// 	})
// }
