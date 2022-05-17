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

func NewTransactionService(repo transactionDomain.Repository) transactionDomain.Service {
	return transactionService{
		Repository: repo,
	}
}

// Ongkir implements transactionDomain.Service
func (ts transactionService) Ongkir(origin int, dest int, weight int, courier string, paket string) (int, string, error) {
	data, err := ts.Repository.Ongkir(origin, dest, weight, courier)

	fmt.Println("ongkir :", data)
	fmt.Println("destination :", courier, paket)
	// fmt.Println("service :", data.Rajaongkir.Results[0].Costs[0].Service)
	if err != nil {
		return 0, "", err
	}
	for i := 0; i < len(data.Rajaongkir.Results[0].Costs)-1; i++ {
		results := data.Rajaongkir.Results[0].Costs
		fmt.Println(results[0].Cost[0].Value)
		if data.Rajaongkir.Results[0].Code == courier && results[i].Service == paket {
			cost := int(results[i].Cost[0].Value)
			etd := results[i].Cost[0].Etd
			fmt.Println("cost ", cost)
			return cost, etd, nil
		}
	}
	return 10, "", errors.New("ongkir can't process")
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
func (ts transactionService) CreateCheckout(code string, domainCheckout transactionDomain.Checkout, domain productDomain.Product) (transactionDomain.Checkout, error) {
	destination, _ := ts.Repository.GetCityId(domainCheckout.Destination)
	origin, _ := ts.Repository.GetCityId(domain.Origin)
	weight := domain.Weight * float64(domainCheckout.Qty)
	cekOngkir := ts.Repository.CheckCourier(origin, destination, int(weight), domainCheckout.Courier, domainCheckout.Package)
	fmt.Println("cekOngkir :", cekOngkir)
	if !cekOngkir {
		return transactionDomain.Checkout{}, errors.New("ongkir not found")
	}
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
