package productMysql

import (
	"errors"

	"github.com/google/uuid"
	productDomain "github.com/gozzafadillah/product/domain"
	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

// GetCategoryById implements productDomain.Repository
func (pr ProductRepo) GetCategoryById(id int) (productDomain.Category, error) {
	var rec Category
	err := pr.DB.Where("id = ?", id).First(&rec).RowsAffected
	if err == 0 {
		return productDomain.Category{}, errors.New("data record not found")
	}
	return toDomainCategory(rec), nil
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

	return products, err
}

// Delete implements productDomain.Repository
func (pr ProductRepo) Delete(id int) error {
	var rec Products
	err := pr.DB.Unscoped().Delete(&rec, id).Error
	return err
}

// Save implements productDomain.Repository
func (pr ProductRepo) SaveProduct(domain productDomain.Product) (int, error) {
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

// GetByNameCategory implements productDomain.Repository
func (pr ProductRepo) GetByNameCategory(name string) ([]productDomain.Product, error) {
	var recCategory Category
	recProduct := []Products{}
	var products []productDomain.Product

	pr.DB.Where("name = ?", name).First(&recCategory)
	err := pr.DB.Where("category_id = ?", recCategory.ID).Find(&recProduct).Error
	for _, value := range recProduct {
		products = append(products, toDomain(value))
	}
	return products, err
}

// SaveCategory implements productDomain.Repository
func (pr ProductRepo) SaveCategory(domain productDomain.Category) (int, error) {
	err := pr.DB.Create(&domain).Error
	return domain.ID, err
}
