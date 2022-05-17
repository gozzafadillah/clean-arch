package request

import (
	"time"

	productDomain "github.com/gozzafadillah/product/domain"
)

type RequestJSON struct {
	ID          int
	Code        string
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Origin      string  `json:"origin"`
	Qty         int     `json:"qty"`
	Price       int     `json:"price"`
	Weight      float64 `json:"weight"`
	Status      bool    `json:"status"`
	Category_Id int     `json:"category_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RequestJSONCategory struct {
	ID        int
	Name      string `json:"name"`
	Status    bool   `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToDomainCategory(req RequestJSONCategory) productDomain.Category {
	return productDomain.Category{
		Name:   req.Name,
		Status: req.Status,
	}
}

func ToDomain(req RequestJSON) productDomain.Product {
	return productDomain.Product{
		Name:        req.Name,
		Description: req.Description,
		Origin:      req.Origin,
		Qty:         req.Qty,
		Price:       req.Price,
		Weight:      req.Weight,
		Status:      req.Status,
		Category_Id: req.Category_Id,
	}
}
