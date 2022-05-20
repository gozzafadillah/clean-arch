package transactionDomain

import productDomain "github.com/gozzafadillah/product/domain"

// import productDomain "github.com/gozzafadillah/product/domain"

type Service interface {
	KalkulationWeight(productWeight int, requestQty int) int
	GetCode() string
	// Transaction
	CreateTransaction(idUser int, code string, ongkir int, etd string, checkout Checkout) (Transaction, error)
	DestroyTransaction(id int) error

	// Checkout
	CreateCheckout(code string, domainCheckout Checkout, domain productDomain.Product, ongkir int, etd string) (Checkout, error)
	GetCheckout(code string) (Checkout, error)
	ChangeStatus(code string) error

	// update qty stok produk
	UpdateStok(id int, qty int) error

	// third party
	CheckCity(city string) (int, error)
	Ongkir(origin, dest int, weight int, courier string, paket string) (int, string, error)
}

type Repository interface {
	// Save / Create data
	SaveTransaction(code string, idUser int, ongkir int, etd string, checkout Checkout) (int, error)
	SaveCheckout(domainCheckout Checkout, domain productDomain.Product, code string, ongkir int, etd string) (int, error)
	// Get data
	GetCheckoutCode(code string) (Checkout, error)
	GetCheckoutId(id int) (Checkout, error)
	GetCode(code string) (Transaction, error)
	GetTransaction(id int) (Transaction, error)
	GetCityId(name string) (int, error)

	CheckCourier(origin int, cityDest int, weight int, courier string, paket string) bool
	Ongkir(origin int, cityDest int, weight int, courier string) (Ongkir, error)
	UpdateQty(id int, qty int) error
	ChangeStatus(code string) error
}
