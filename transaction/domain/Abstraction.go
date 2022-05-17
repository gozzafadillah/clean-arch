package transactionDomain

import productDomain "github.com/gozzafadillah/product/domain"

// import productDomain "github.com/gozzafadillah/product/domain"

type Service interface {
	GetCode() string
	// Transaction
	CreateTransaction(idUser int, code string, ongkir int, etd string, checkout Checkout) (Transaction, error)
	DestroyTransaction(id int) error

	// Checkout
	CreateCheckout(code string, domainCheckout Checkout, domain productDomain.Product) (Checkout, error)

	// update qty stok produk
	UpdateStok(id int, qty int) error

	// third party
	CheckCity(city string) (int, error)
	Ongkir(origin, dest int, weight int, courier string, paket string) (int, string, error)
}

type Repository interface {
	// Save / Create data
	SaveTransaction(code string, idUser int, ongkir int, etd string, checkout Checkout) (int, error)
	SaveCheckout(domainCheckout Checkout, domain productDomain.Product, code string) (int, error)
	// Get data
	GetCheckoutId(id int) (Checkout, error)
	GetCode(code string) (Transaction, error)
	GetTransaction(id int) (Transaction, error)
	GetCityId(name string) (int, error)

	Delete(id int) error
	CheckCourier(origin int, cityDest int, weight int, courier string, paket string) bool
	Ongkir(origin int, cityDest int, weight int, courier string) (Ongkir, error)
	UpdateQty(id int, qty int) error
}
