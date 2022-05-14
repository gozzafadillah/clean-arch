package response

import (
	transactionDomain "github.com/gozzafadillah/transaction/domain"
)

type ResponseJSONTransaction struct {
	Code           string  `json:"code"`
	User_Id        int     `json:"user_id"`
	Total_Qty      int     `json:"total_qty"`
	Total_Price    float64 `json:"total_price"`
	Shipping_Price float64 `json:"shipping_price"`
}

func ToDomainCheckout(req ResponseJSONTransaction) transactionDomain.Transaction {
	return transactionDomain.Transaction{
		Code:           req.Code,
		User_Id:        req.User_Id,
		Total_Qty:      req.Total_Qty,
		Total_Price:    req.Total_Price,
		Shipping_Price: req.Shipping_Price,
	}
}
