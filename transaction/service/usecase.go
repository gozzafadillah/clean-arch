package service

import (
	"errors"
	"fmt"

	productDomain "github.com/gozzafadillah/product/domain"
	transactionDomain "github.com/gozzafadillah/transaction/domain"
	"github.com/pborman/uuid"
)

type transactionService struct {
	Repository transactionDomain.Repository
}

// ChangeStatus implements transactionDomain.Service
func (ts transactionService) ChangeStatus(code string) error {
	err := ts.Repository.ChangeStatus(code)
	if err != nil {
		return errors.New("checkout cannot to change status")
	}
	return nil
}

// KalkulationWeight implements transactionDomain.Service
func (transactionService) KalkulationWeight(productWeight int, requestQty int) int {
	kalkulation := productWeight * requestQty
	return kalkulation
}

// GetCheckout implements transactionDomain.Service
func (ts transactionService) GetCheckout(code string) (transactionDomain.Checkout, error) {
	respon, err := ts.Repository.GetCheckoutCode(code)
	if err != nil {
		return transactionDomain.Checkout{}, errors.New("checkout not found")
	}
	return respon, nil
}

func NewTransactionService(repo transactionDomain.Repository) transactionDomain.Service {
	return transactionService{
		Repository: repo,
	}
}

// UpdateStok implements transactionDomain.Service
func (ts transactionService) UpdateStok(id int, qty int) error {
	err := ts.Repository.UpdateQty(id, qty)
	if err != nil {
		return errors.New("qty cant update")
	}
	return nil
}

// Ongkir implements transactionDomain.Service
func (ts transactionService) Ongkir(origin int, dest int, weight int, courier string, paket string) (int, string, error) {
	cekOngkir := ts.Repository.CheckCourier(origin, dest, weight, courier, paket)
	fmt.Println("data :", origin, dest, weight, courier, paket)
	fmt.Println("status cek ongkir :", cekOngkir)
	if !cekOngkir {
		return 0, "", errors.New("ongkir not found")
	}
	data, err := ts.Repository.Ongkir(origin, dest, weight, courier)
	if err != nil {
		return 0, "", err
	}
	for i := 0; i < len(data.Rajaongkir.Results[0].Costs)-1; i++ {
		results := data.Rajaongkir.Results[0].Costs
		if data.Rajaongkir.Results[0].Code == courier && results[i].Service == paket {
			cost := int(results[i].Cost[0].Value)
			etd := results[i].Cost[0].Etd
			return cost, etd, nil
		}
	}
	return 0, "", errors.New("ongkir can't process")
}

// CheckCity implements transactionDomain.Service
func (ts transactionService) CheckCity(city string) (int, error) {
	data, err := ts.Repository.GetCityId(city)

	if err != nil {
		return 0, errors.New("kota tidak ditemukan")
	}

	return data, nil
}

// CreateCheckout implements transactionDomain.Service
func (ts transactionService) CreateCheckout(code string, domainCheckout transactionDomain.Checkout, domain productDomain.Product, ongkir int, etd string) (transactionDomain.Checkout, error) {
	id, err := ts.Repository.SaveCheckout(domainCheckout, domain, code, ongkir, etd)
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
func (ts transactionService) CreateTransaction(idUser int, code string, ongkir int, etd string, checkout transactionDomain.Checkout) (transactionDomain.Transaction, error) {
	id, err := ts.Repository.SaveTransaction(code, idUser, ongkir, etd, checkout)
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
func (ts transactionService) DestroyTransaction(id int) error {
	panic("no nono")
}
