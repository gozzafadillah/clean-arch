package response

import (
	transactionDomain "github.com/gozzafadillah/transaction/domain"
)

type ResponseJSONTransaction struct {
	Code             string  `json:"code"`
	User_Id          int     `json:"user_id"`
	Total_Qty        int     `json:"total_qty"`
	Total_Price      float64 `json:"total_price"`
	Shipping_Name    string  `json:"shipping_name"`
	Shipping_Package string  `json:"shipping_package"`
	Etd              string  `json:"estimate"`
}

type ResponseJSONCheckout struct {
	TransactionCode string  `json:"transaction_code"`
	Qty             int     `json:"qty"`
	Price           float64 `json:"price"`
	Weight          float64 `json:"weight"`
	Destination     string  `json:"destination"`
	Courier         string  `json:"courier"`
	Package         string  `json:"package"`
	Etd             string  `json:"etimate"`
	Shipping_Price  float64 `json:"shipping_price"`
	Status          bool    `json:"status"`
}

func FromDomainTransaction(domain transactionDomain.Transaction) ResponseJSONTransaction {
	return ResponseJSONTransaction{
		Code:             domain.Code,
		User_Id:          domain.User_Id,
		Total_Qty:        domain.Total_Qty,
		Total_Price:      domain.Total_Price,
		Shipping_Name:    domain.Shipping_Name,
		Shipping_Package: domain.Shipping_Package,
		Etd:              domain.Etd,
	}
}
func FromDomainCheckout(domain transactionDomain.Checkout) ResponseJSONCheckout {
	return ResponseJSONCheckout{
		TransactionCode: domain.TransactionCode,
		Qty:             domain.Qty,
		Price:           domain.Price,
		Weight:          domain.Weight,
		Destination:     domain.Destination,
		Courier:         domain.Courier,
		Package:         domain.Package,
		Etd:             domain.Etd,
		Shipping_Price:  domain.Shipping_Price,
		Status:          domain.Status,
	}
}
