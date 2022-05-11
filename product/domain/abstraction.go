package productDomain

type Service interface {
	// CRUD Product
	GetProducts() ([]Product, error)
	GetProductId(id int) (Product, error)
	CreateProduct(domain Product) (Product, error)
	EditProduct(id int, domain Product) (Product, error)
	DestroyProduct(id int) (Product, error)
	// CRUD Category
	CreateCategory(domain Category) (Category, error)
	GetCategoryById(id int) (Category, error)

	// feature
	GetMinPrice() ([]Product, error)
	GetMaxPrice() ([]Product, error)
	GetCategory(name string) ([]Product, error)
}

type Repository interface {
	GetProducts() ([]Product, error)
	GetById(id int) (Product, error)
	GetCategoryById(id int) (Category, error)
	GetByNameCategory(name string) ([]Product, error)
	SaveProduct(domain Product) (int, error)
	SaveCategory(domain Category) (int, error)
	Update(id int, domain Product) error
	Delete(id int) error
}
