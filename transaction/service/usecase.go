package service

import (
	"errors"
	"fmt"
	"strconv"

	productDomain "github.com/gozzafadillah/product/domain"
	transactionDomain "github.com/gozzafadillah/transaction/domain"
	"github.com/pborman/uuid"
)

type transactionService struct {
	Repository transactionDomain.Repository
}

// CheckCity implements transactionDomain.Service
func (ts transactionService) CheckCity(city string) (int, error) {
	data, err := ts.Repository.GetCityId()
	var id int
	if err != nil {
		return 0, errors.New("kota tidak ditemukan")
	}
	fmt.Println(len(data.Rajaongkir.Results))
	for i := 1; i <= len(data.Rajaongkir.Results); i++ {
		if data.Rajaongkir.Results[i].CityName == city {
			cityId, _ := strconv.Atoi(data.Rajaongkir.Results[i].CityID)
			fmt.Println("city id :", cityId)
			id = cityId
			return id, nil
		}
	}
	return 0, err
}

// Ongkir implements transactionDomain.Service
func (transactionService) Ongkir(origin int, cityDest int) (transactionDomain.Ongkir, error) {
	panic("unimplemented")
}

// CreateCheckout implements transactionDomain.Service
func (ts transactionService) CreateCheckout(code string, domainCheckout transactionDomain.Checkout, domain productDomain.Product) (transactionDomain.Checkout, error) {
	id, err := ts.Repository.SaveCheckout(domainCheckout, domain, code)
	fmt.Println("id : ", id)
	if err != nil {
		return transactionDomain.Checkout{}, errors.New("bad request")
	}
	dataCheckout, err := ts.Repository.GetCheckoutId(id)
	if err != nil {
		return transactionDomain.Checkout{}, errors.New("checkout not found, bad request")
	}
	return dataCheckout, nil
}

// GetCode implements transactionDomain.Service
func (ts transactionService) GetCode() string {
	generateCode := uuid.NewRandom()
	return generateCode.String()
}

// CreateTransaction implements transactionDomain.Service
func (ts transactionService) CreateTransaction(idUser int, code string, checkout transactionDomain.Checkout) (transactionDomain.Transaction, error) {
	id, err := ts.Repository.SaveTransaction(code, idUser, checkout)
	if err != nil {
		return transactionDomain.Transaction{}, errors.New("data not found")
	}
	dataTransaction, err := ts.Repository.GetTransaction(id)
	if err != nil {
		return transactionDomain.Transaction{}, errors.New("data not found")
	}
	return dataTransaction, nil
}

// DestroyTransaction implements transactionDomain.Service
func (transactionService) DestroyTransaction(id int) error {
	panic("unimplemented")
}

func NewTransactionService(repo transactionDomain.Repository) transactionDomain.Service {
	return transactionService{
		Repository: repo,
	}
}
