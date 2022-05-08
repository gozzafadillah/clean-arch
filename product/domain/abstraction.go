package productDomain

type Service interface {
	GetProducts() ([]Product, error)
	GetProductId(id int) (Product, error)
	GetMinPrice(domain Product) ([]Product, error)
	CreateProduct(domain Product) (Product, error)
	EditProduct(id int, domain Product) (Product, error)
	DestroyProduct(id int) (Product, error)
}

type Repository interface {
	GetProducts() ([]Product, error)
	GetById(id int) (Product, error)
	Save(domain Product) (int, error)
	Update(id int, domain Product) error
	Delete(id int) error
}
