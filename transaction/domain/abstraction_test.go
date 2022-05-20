package transactionDomain_test

import (
	"os"
	"testing"

	productDomain "github.com/gozzafadillah/product/domain"
	transactionDomain "github.com/gozzafadillah/transaction/domain"
	transactionMock "github.com/gozzafadillah/transaction/domain/mocks"
	serviceTransaction "github.com/gozzafadillah/transaction/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	transactionService transactionDomain.Service
	domainTransaction  transactionDomain.Transaction
	domainCheckout     transactionDomain.Checkout
	domainOngkir       transactionDomain.Ongkir
	domainCity         transactionDomain.CityRO
	domainProduct      productDomain.Product
	transactionRepo    transactionMock.Repository
)

func TestMain(m *testing.M) {
	transactionService = serviceTransaction.NewTransactionService(&transactionRepo)
	domainOngkir = transactionDomain.Ongkir{}
	domainCity = transactionDomain.CityRO{}
	domainTransaction = transactionDomain.Transaction{
		ID:               1,
		Code:             "abcd-efgh-ijkl",
		User_Id:          1,
		Total_Qty:        2,
		Total_Price:      50000,
		Shipping_Name:    "jne",
		Shipping_Package: "OKE",
		Shipping_Price:   15000,
		Etd:              "2-3",
	}
	domainCheckout = transactionDomain.Checkout{
		Qty:         2,
		Destination: "Bandung",
		Courier:     "jne",
		Package:     "OKE",
	}
	domainProduct = productDomain.Product{
		ID:          1,
		Code:        "abcd-efgh-ijkl",
		Name:        "product_1",
		Description: "description product 1",
		Origin:      "Bandung",
		Qty:         5,
		Price:       5000,
		Weight:      2000,
		Status:      true,
		Category_Id: 1,
	}
	os.Exit(m.Run())
}

func TestCheckCity(t *testing.T) {
	t.Run("check city", func(t *testing.T) {
		transactionRepo.On("GetCityId", mock.Anything).Return(22, nil).Once()

		res, err := transactionService.CheckCity("Bandung")

		assert.NoError(t, err)
		assert.Equal(t, 22, res)
	})
}

func TestCreateCheckout(t *testing.T) {
	t.Run("create checkout", func(t *testing.T) {
		transactionRepo.On("SaveCheckout", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(1, nil).Once()
		transactionRepo.On("GetCheckoutId", mock.Anything).Return(domainCheckout, nil).Once()

		res, err := transactionService.CreateCheckout("abcd-efgh-ijkl", domainCheckout, domainProduct, 1500, "2-3")

		assert.NoError(t, err)
		assert.Equal(t, "Bandung", res.Destination)
	})
}

func TestCreateTransaction(t *testing.T) {
	t.Run("create transaction", func(t *testing.T) {
		transactionRepo.On("SaveTransaction", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(1, nil).Once()
		transactionRepo.On("GetTransaction", mock.Anything).Return(domainTransaction, nil).Once()

		res, err := transactionService.CreateTransaction(1, "abcd-efgh-ijkl", 15000, "2-3", domainCheckout)

		assert.NoError(t, err)
		assert.Equal(t, "abcd-efgh-ijkl", res.Code)
	})
}

func TestUpdateStock(t *testing.T) {
	t.Run("update stock product", func(t *testing.T) {
		transactionRepo.On("UpdateQty", mock.Anything, mock.Anything).Return(nil).Once()

		err := transactionService.UpdateStok(1, 3)

		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
}
func TestChangeStatus(t *testing.T) {
	t.Run("change status checkout", func(t *testing.T) {
		transactionRepo.On("ChangeStatus", mock.Anything).Return(nil).Once()

		err := transactionService.ChangeStatus("abcd-efgh-ijkl")

		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
}

func TestGetCheckout(t *testing.T) {
	t.Run("get checkout by code", func(t *testing.T) {
		domainCheckout.TransactionCode = "abcd-efgh-ijkl"
		transactionRepo.On("GetCheckoutCode", mock.Anything).Return(domainCheckout, nil).Once()

		res, err := transactionService.GetCheckout("abcd-efgh-ijkl")

		assert.NoError(t, err)
		assert.Equal(t, "abcd-efgh-ijkl", res.TransactionCode)
	})
}
