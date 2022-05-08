package response

import productDomain "github.com/gozzafadillah/product/domain"

type ResponseJSON struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Qty         int     `json:"qty"`
	Price       int     `json:"price"`
	Weight      float64 `json:"weight"`
	Status      bool    `json:"status"`
	Category_Id int     `json:"category_id"`
}

func FromDomain(domain productDomain.Product) ResponseJSON {
	return ResponseJSON{
		Code:        domain.Code,
		Name:        domain.Name,
		Description: domain.Description,
		Qty:         domain.Qty,
		Price:       domain.Price,
		Weight:      domain.Weight,
		Status:      domain.Status,
		Category_Id: domain.Category_Id,
	}
}
