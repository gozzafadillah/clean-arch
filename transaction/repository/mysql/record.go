package mysql

import (
	"time"

	transactionDomain "github.com/gozzafadillah/transaction/domain"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID               int
	Code             string
	User_Id          int
	Total_Qty        int
	Total_Price      float64
	Shipping_Price   float64
	Shipping_Name    string
	Shipping_Package string
	Etd              string
}
type Checkout struct {
	gorm.Model
	ID              int
	ProductID       int
	TransactionCode string
	Destination     string
	Qty             int
	Price           float64
	Weight          float64
	Courier         string
	Package         string
	Etd             string
	Shipping_Price  float64
	Status          bool
}

func toDomain(rec Transaction) transactionDomain.Transaction {
	return transactionDomain.Transaction{
		ID:               rec.ID,
		Code:             rec.Code,
		User_Id:          rec.User_Id,
		Total_Qty:        rec.Total_Qty,
		Total_Price:      rec.Total_Price,
		Shipping_Name:    rec.Shipping_Name,
		Shipping_Package: rec.Shipping_Package,
		Shipping_Price:   rec.Shipping_Price,
		Etd:              rec.Etd,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	}
}

func toDomainCheckout(rec Checkout) transactionDomain.Checkout {
	return transactionDomain.Checkout{
		ID:              rec.ID,
		ProductID:       rec.ProductID,
		TransactionCode: rec.TransactionCode,
		Qty:             rec.Qty,
		Price:           rec.Price,
		Weight:          rec.Weight,
		Destination:     rec.Destination,
		Courier:         rec.Courier,
		Package:         rec.Package,
		Etd:             rec.Etd,
		Shipping_Price:  rec.Shipping_Price,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
	}
}
