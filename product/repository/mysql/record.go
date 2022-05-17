package productMysql

import (
	"time"

	productDomain "github.com/gozzafadillah/product/domain"
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ID          int
	Code        string
	Name        string
	Origin      string
	Category_Id int
	Description string
	Qty         int
	Price       int
	Weight      float64
	Status      bool
}

type Category struct {
	gorm.Model
	ID     int
	Name   string
	Status bool
}

// Category
func toDomainCategory(rec Category) productDomain.Category {
	return productDomain.Category{
		ID:     rec.ID,
		Name:   rec.Name,
		Status: rec.Status,
	}
}

func fromDomainCategory(rec productDomain.Category) Category {
	return Category{
		ID:     rec.ID,
		Name:   rec.Name,
		Status: rec.Status,
	}
}

// Product
func toDomain(rec Products) productDomain.Product {
	return productDomain.Product{
		ID:          rec.ID,
		Code:        rec.Code,
		Name:        rec.Name,
		Description: rec.Description,
		Origin:      rec.Origin,
		Qty:         rec.Qty,
		Price:       rec.Price,
		Weight:      rec.Weight,
		Status:      rec.Status,
		Category_Id: rec.Category_Id,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func fromDomain(rec productDomain.Product) Products {
	return Products{
		ID:          rec.ID,
		Code:        rec.Code,
		Name:        rec.Name,
		Origin:      rec.Origin,
		Category_Id: rec.Category_Id,
		Description: rec.Description,
		Qty:         rec.Qty,
		Price:       rec.Price,
		Weight:      rec.Weight,
		Status:      rec.Status,
	}
}
