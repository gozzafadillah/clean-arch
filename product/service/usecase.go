package serviceProduct

import (
	"errors"

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

// CheckoutProductId implements productDomain.Service
func (ps ProductService) CheckoutProductId(id int) (productDomain.Product, error) {
	data, err := ps.Repository.GetById(id)
	if err != nil {
		return productDomain.Product{}, errors.New("data tidak adaa")
	}
	return data, nil
}

// GetProducts implements productDomain.Service
func (ps ProductService) GetProducts() ([]productDomain.Product, error) {
	data, err := ps.Repository.GetProducts()
	if err != nil {
		return []productDomain.Product{}, errors.New("data tidak adaa")
	}
	return data, nil
}

// GetProductId implements productDomain.Service
func (ps ProductService) GetProductId(id int) (productDomain.Product, error) {
	data, err := ps.Repository.GetById(id)
	if err != nil {
		return productDomain.Product{}, errors.New("data tidak adaa")
	}
	return data, nil
}

// CreateProduct implements productDomain.Service
func (ps ProductService) CreateProduct(domain productDomain.Product) (productDomain.Product, error) {
	id, err := ps.Repository.SaveProduct(domain)
	if err != nil {
		return productDomain.Product{}, errors.New("bad request")
	}
	data, err := ps.Repository.GetById(id)
	if err != nil {
		return productDomain.Product{}, errors.New("data tidak adaa")
	}

	return data, err
}

// DestroyProduct implements productDomain.Service
func (ps ProductService) DestroyProduct(id int) (productDomain.Product, error) {
	data, err := ps.Repository.GetById(id)
	if err != nil {
		return productDomain.Product{}, errors.New("not found")
	}
	delete := ps.Repository.Destroy(id)
	if err := delete; err != nil {
		return productDomain.Product{}, errors.New("not found")
	}

	return data, nil
}

// EditProduct implements productDomain.Service
func (ps ProductService) EditProduct(id int, domain productDomain.Product) (productDomain.Product, error) {
	err := ps.Repository.Update(id, domain)
	if err != nil {
		return productDomain.Product{}, errors.New("data didn't update")
	}
	data, err := ps.Repository.GetById(id)
	if err != nil {
		return productDomain.Product{}, errors.New("data unknown")
	}
	return data, nil
}

// GetMinPrice implements productDomain.Service
func SortMin(data []productDomain.Product) {
	var isDone = false
	// Sort asc
	for !isDone {
		isDone = true
		var i = 0
		for i < len(data)-1 {
			if data[i].Price > data[i+1].Price {
				data[i], data[i+1] = data[i+1], data[i]
				isDone = false
			}
			i++
		}
	}
}

func (ps ProductService) GetMinPrice() ([]productDomain.Product, error) {
	data, err := ps.Repository.GetProducts()
	if err != nil {
		return []productDomain.Product{}, errors.New("data empty")
	}

	SortMin(data)

	return data, nil
}

// GetMaxPrice implements productDomain.Service
func SortMAX(data []productDomain.Product) {
	var isDone = false
	// Sort Desc
	for !isDone {
		isDone = true
		var i = 0
		for i < len(data)-1 {
			if data[i].Price < data[i+1].Price {
				data[i], data[i+1] = data[i+1], data[i]
				isDone = false
			}
			i++
		}
	}
}
func (ps ProductService) GetMaxPrice() ([]productDomain.Product, error) {
	data, err := ps.Repository.GetProducts()
	if err != nil {
		return []productDomain.Product{}, errors.New("data empty")
	}
	SortMAX(data)

	return data, nil
}

// GetCategory implements productDomain.Service
func (ps ProductService) GetCategory(name string) ([]productDomain.Product, error) {
	data, err := ps.Repository.GetByNameCategory(name)
	if err != nil {
		return []productDomain.Product{}, errors.New("category not found")
	}
	return data, nil
}

// GetCategoryById implements productDomain.Service
func (ps ProductService) GetCategoryById(id int) (productDomain.Category, error) {
	data, err := ps.Repository.GetCategoryById(id)
	if err != nil {
		return productDomain.Category{}, errors.New("data tidak adaa")
	}
	return data, nil
}

// CreateCategory implements productDomain.Service
func (ps ProductService) CreateCategory(domain productDomain.Category) (productDomain.Category, error) {
	id, err := ps.Repository.SaveCategory(domain)
	if err != nil {
		return productDomain.Category{}, errors.New("category not save, check again")
	}
	data, err := ps.Repository.GetCategoryById(id)
	if err != nil {
		return productDomain.Category{}, errors.New("category not found")
	}

	return data, err
}
