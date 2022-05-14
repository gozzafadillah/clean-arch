package transactionDomain

import productDomain "github.com/gozzafadillah/product/domain"

// import productDomain "github.com/gozzafadillah/product/domain"

type Service interface {
	GetCode() string
	// Transaction
	CreateTransaction(idUser int, code string, checkout Checkout) (Transaction, error)
	DestroyTransaction(id int) error

	// Checkout
	CreateCheckout(code string, domainCheckout Checkout, domain productDomain.Product) (Checkout, error)

	// third party
	CheckCity(city string) (int, error)
	Ongkir(origin int, cityDest int) (Ongkir, error)
}

type Repository interface {
	// Save / Create data
	SaveTransaction(code string, idUser int, checkout Checkout) (int, error)
	SaveCheckout(domainCheckout Checkout, domain productDomain.Product, code string) (int, error)
	// Get data
	GetCheckoutId(id int) (Checkout, error)
	GetCode(code string) (Transaction, error)
	GetTransaction(id int) (Transaction, error)
	GetCityId() (CityRO, error)

	Delete(id int) error
}
