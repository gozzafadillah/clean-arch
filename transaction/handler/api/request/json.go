package request

import (
	transactionDomain "github.com/gozzafadillah/transaction/domain"
)

type RequestJSONCheckout struct {
	ID              int
	ProductID       int
	TransactionCode string
	Qty             int `json:"qty" form:"qty"`
	Price           float64
	Weight          float64
	Destination     string `json:"destination" form:"destination"`
	Courier         string `json:"courier" form:"courier"`
	Package         string `json:"package" form:"package"`
	Status          bool
}

func ToDomainCheckout(req RequestJSONCheckout) transactionDomain.Checkout {
	return transactionDomain.Checkout{
		Qty:         req.Qty,
		Destination: req.Destination,
		Courier:     req.Courier,
		Package:     req.Package,
	}
}
