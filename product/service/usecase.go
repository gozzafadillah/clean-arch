package service

import (
	"errors"
	"fmt"

	errorConv "github.com/gozzafadillah/helper/error"
	productDomain "github.com/gozzafadillah/product/domain"
)

type ProductService struct {
	Repository productDomain.Repository
}

func NewProductService(repo productDomain.Repository) productDomain.Service {
	return ProductService{
		Repository: repo,
	}
}

// GetProducts implements productDomain.Service
func (ps ProductService) GetProducts() ([]productDomain.Product, error) {
	data, err := ps.Repository.GetProducts()
	if err != nil {
		return []productDomain.Product{}, errorConv.Conversion(err)
	}
	return data, nil
}

// GetProductId implements productDomain.Service
func (ps ProductService) GetProductId(id int) (productDomain.Product, error) {
	data, err := ps.Repository.GetById(id)
	if err != nil {
		return productDomain.Product{}, errorConv.Conversion(err)
	}
	return data, nil
}

// CreateProduct implements productDomain.Service
func (ps ProductService) CreateProduct(domain productDomain.Product) (productDomain.Product, error) {
	id, err := ps.Repository.Save(domain)
	fmt.Println("id : ", id)
	if err != nil {
		return productDomain.Product{}, errorConv.Conversion(err)
	}
	data, err := ps.Repository.GetById(id)
	if err != nil {
		return productDomain.Product{}, errorConv.Conversion(err)
	}

	return data, err
}

// DestroyProduct implements productDomain.Service
func (ps ProductService) DestroyProduct(id int) (productDomain.Product, error) {
	data, err1 := ps.Repository.GetById(id)
	if err1 != nil {
		return productDomain.Product{}, errors.New("not found")
	}
	err2 := ps.Repository.Delete(id)
	fmt.Println("id : ", id)
	if err2 != nil {
		errorConv.Conversion(err2)
	}

	return data, nil
}

// EditProduct implements productDomain.Service
func (ps ProductService) EditProduct(id int, domain productDomain.Product) (productDomain.Product, error) {
	err := ps.Repository.Update(id, domain)
	if err != nil {
		return productDomain.Product{}, errorConv.Conversion(err)
	}
	data, err := ps.Repository.GetById(id)
	if err != nil {
		return productDomain.Product{}, errorConv.Conversion(err)
	}
	return data, nil
}

// GetMinPrice implements productDomain.Service
func (ProductService) GetMinPrice(domain productDomain.Product) ([]productDomain.Product, error) {
	panic("unimplemented")
}
