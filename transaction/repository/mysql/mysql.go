package mysql

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gozzafadillah/app/config"
	productDomain "github.com/gozzafadillah/product/domain"
	transactionDomain "github.com/gozzafadillah/transaction/domain"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transactionDomain.Repository {
	return TransactionRepo{
		DB: db,
	}
}

// GetCityId implements transactionDomain.Repository
func (ts TransactionRepo) GetCityId() (transactionDomain.CityRO, error) {
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
	return responseObject, nil
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
func (tr TransactionRepo) SaveCheckout(domainCheckout transactionDomain.Checkout, domain productDomain.Product, code string) (int, error) {
	domainCheckout.Price = float64(domain.Price * domainCheckout.Qty)
	domainCheckout.Weight = domain.Weight * float64(domainCheckout.Qty)
	domainCheckout.TransactionCode = code
	domainCheckout.ProductID = domain.ID
	err := tr.DB.Create(&domainCheckout).Error
	return domainCheckout.ID, err
}

// Save implements transactionDomain.Repository
func (tr TransactionRepo) SaveTransaction(code string, idUser int, checkout transactionDomain.Checkout) (int, error) {
	var transaction transactionDomain.Transaction
	transaction.Shipping_Price = float64(3300) * checkout.Weight
	transaction.Total_Price = checkout.Price + transaction.Shipping_Price
	transaction.Total_Qty = checkout.Qty
	transaction.Code = code
	transaction.User_Id = idUser
	transaction.Shipping_Id = 1
	err := tr.DB.Create(&transaction).Error
	return transaction.ID, err
}

// Delete implements transactionDomain.Repository
func (tr TransactionRepo) Delete(id int) error {
	rec := Transaction{}
	err := tr.DB.Unscoped().Delete(&rec, id).Error
	return err
}
