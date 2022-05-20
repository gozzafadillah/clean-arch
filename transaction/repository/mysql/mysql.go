package mysql

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gozzafadillah/app/config"
	productDomain "github.com/gozzafadillah/product/domain"
	transactionDomain "github.com/gozzafadillah/transaction/domain"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

// ChangeStatus implements transactionDomain.Repository
func (tr TransactionRepo) ChangeStatus(code string) error {
	rec := Checkout{}
	err := tr.DB.Model(&rec).Where("transaction_code = ?", code).Update("status", true).Error
	return err
}

// GetCheckoutCode implements transactionDomain.Repository
func (tr TransactionRepo) GetCheckoutCode(code string) (transactionDomain.Checkout, error) {
	rec := Checkout{}
	err := tr.DB.Where("transaction_code", code).First(&rec).Error
	return toDomainCheckout(rec), err
}

func NewTransactionRepository(db *gorm.DB) transactionDomain.Repository {
	return TransactionRepo{
		DB: db,
	}
}

// UpdateQty implements transactionDomain.Repository
func (tr TransactionRepo) UpdateQty(id int, qty int) error {
	rec := productDomain.Product{}
	err := tr.DB.Model(&rec).Where("id = ?", id).Update("qty", qty).Error
	return err
}

// CheckCourier implements transactionDomain.Repository
func (TransactionRepo) CheckCourier(origin int, cityDest int, weight int, courier string, paket string) bool {
	url := config.BaseURLRO + "cost"
	client := &http.Client{}
	payload := strings.NewReader("origin=" + strconv.Itoa(origin) + "&destination=" + strconv.Itoa(cityDest) + "&weight=" + strconv.Itoa(weight) + "&courier=" + courier)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("key", config.Key)

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject transactionDomain.Ongkir
	json.Unmarshal(body, &responseObject)
	for i := 0; i < len(responseObject.Rajaongkir.Results); i++ {
		if responseObject.Rajaongkir.Results[0].Code == courier && responseObject.Rajaongkir.Results[0].Costs[i].Service == paket {
			return true
		}
	}
	return false
}

// Ongkir implements transactionDomain.Repository
func (tr TransactionRepo) Ongkir(origin int, cityDest int, weight int, courier string) (transactionDomain.Ongkir, error) {
	url := config.BaseURLRO + "cost"
	client := &http.Client{}
	payload := strings.NewReader("origin=" + strconv.Itoa(origin) + "&destination=" + strconv.Itoa(cityDest) + "&weight=" + strconv.Itoa(weight) + "&courier=" + courier)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("key", config.Key)

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject transactionDomain.Ongkir
	json.Unmarshal(body, &responseObject)
	return responseObject, nil
}

// GetCityId implements transactionDomain.Repository
func (ts TransactionRepo) GetCityId(city string) (int, error) {
	url := config.BaseURLRO + "city"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", config.Key)

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var responseObject transactionDomain.CityRO
	json.Unmarshal(body, &responseObject)

	data := responseObject
	var id int
	for i := 0; i < len(data.Rajaongkir.Results)-1; i++ {
		if data.Rajaongkir.Results[i].CityName == city {
			cityId, _ := strconv.Atoi(data.Rajaongkir.Results[i].CityID)
			fmt.Println("city id :", cityId)
			id = cityId
			return id, nil
		}
	}
	return id, errors.New("city not found")
}

// GetTransaction implements transactionDomain.Repository
func (tr TransactionRepo) GetTransaction(id int) (transactionDomain.Transaction, error) {
	var rec Transaction
	err := tr.DB.Where("id = ?", id).First(&rec).Error
	return toDomain(rec), err
}

// GetCheckoutId implements transactionDomain.Repository
func (tr TransactionRepo) GetCheckoutId(id int) (transactionDomain.Checkout, error) {
	var rec Checkout
	err := tr.DB.Where("id = ?", id).First(&rec).Error
	return toDomainCheckout(rec), err
}

// GetCode implements transactionDomain.Repository
func (tr TransactionRepo) GetCode(code string) (transactionDomain.Transaction, error) {
	var rec Transaction
	err := tr.DB.Where("code = ?", code).First(&rec).Error
	return toDomain(rec), err
}

// SaveCheckout implements transactionDomain.Repository
func (tr TransactionRepo) SaveCheckout(domainCheckout transactionDomain.Checkout, domain productDomain.Product, code string, ongkir int, etd string) (int, error) {
	domainCheckout.Price = float64(domain.Price * domainCheckout.Qty)
	domainCheckout.Weight = domain.Weight * float64(domainCheckout.Qty)
	domainCheckout.TransactionCode = code
	domainCheckout.ProductID = domain.ID
	domainCheckout.Shipping_Price = float64(ongkir)
	domainCheckout.Etd = etd
	err := tr.DB.Create(&domainCheckout).Error
	return domainCheckout.ID, err
}

// Save implements transactionDomain.Repository
func (tr TransactionRepo) SaveTransaction(code string, idUser int, ongkir int, etd string, checkout transactionDomain.Checkout) (int, error) {
	var transaction transactionDomain.Transaction
	transaction.Shipping_Price = float64(ongkir)
	transaction.Total_Qty = checkout.Qty
	transaction.Code = code
	transaction.User_Id = idUser
	transaction.Shipping_Name = checkout.Courier
	transaction.Shipping_Package = checkout.Package
	transaction.Total_Price = checkout.Price + transaction.Shipping_Price
	transaction.Etd = etd
	err := tr.DB.Create(&transaction).Error
	return transaction.ID, err
}
