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
	Shipping_Price   float64 `json:"shipping_price"`
	Etd              string  `json:"estimate"`
}

func FromDomainCheckout(domain transactionDomain.Transaction) ResponseJSONTransaction {
	return ResponseJSONTransaction{
		Code:             domain.Code,
		User_Id:          domain.User_Id,
		Total_Qty:        domain.Total_Qty,
		Total_Price:      domain.Total_Price,
		Shipping_Price:   domain.Shipping_Price,
		Shipping_Name:    domain.Shipping_Name,
		Shipping_Package: domain.Shipping_Package,
		Etd:              domain.Etd,
	}
}
