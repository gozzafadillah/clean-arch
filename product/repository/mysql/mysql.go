package productMysql

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	productDomain "github.com/gozzafadillah/product/domain"
	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) productDomain.Repository {
	return ProductRepo{
		DB: db,
	}
}

// GetProducts implements productDomain.Repository
func (pr ProductRepo) GetProducts() ([]productDomain.Product, error) {
	rec := []Products{}
	err := pr.DB.Find(&rec).Error

	var products []productDomain.Product
	for _, value := range rec {
		products = append(products, toDomain(value))
	}
	fmt.Println("array products = ", products)

	return products, err
}

// Delete implements productDomain.Repository
func (pr ProductRepo) Delete(id int) error {
	var rec Products
	err := pr.DB.Unscoped().Delete(&rec, id).Error
	return err
}

// Save implements productDomain.Repository
func (pr ProductRepo) Save(domain productDomain.Product) (int, error) {
	data, _ := uuid.NewRandom()
	domain.Code = data.String()

	err := pr.DB.Create(&domain).Error
	return domain.ID, err
}

// GetById implements productDomain.Repository
func (pr ProductRepo) GetById(id int) (productDomain.Product, error) {
	var rec Products
	err := pr.DB.Where("id = ?", id).First(&rec).RowsAffected
	if err == 0 {
		return productDomain.Product{}, errors.New("data record not found")
	}
	return toDomain(rec), nil
}

// Update implements productDomain.Repository
func (pr ProductRepo) Update(id int, domain productDomain.Product) error {
	newRec := map[string]interface{}{
		"name":        domain.Name,
		"description": domain.Description,
		"qty":         domain.Qty,
		"price":       domain.Price,
		"weight":      domain.Weight,
		"status":      domain.Status,
	}
	err := pr.DB.Model(&domain).Where("id = ?", id).Updates(newRec).Error

	return err
}
