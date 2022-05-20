package productDomain_test

import (
	"os"
	"testing"

	productDomain "github.com/gozzafadillah/product/domain"
	productMock "github.com/gozzafadillah/product/domain/mocks"
	serviceProduct "github.com/gozzafadillah/product/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	productService productDomain.Service
	domainProduct  productDomain.Product
	domainCategory productDomain.Category
	productRepo    productMock.Repository
)

func TestMain(m *testing.M) {
	productService = serviceProduct.NewProductService(&productRepo)
	domainProduct = productDomain.Product{
		ID:          1,
		Code:        "abcd-efgh-ijkl",
		Name:        "product_1",
		Description: "description product 1",
		Origin:      "Bandung",
		Qty:         5,
		Price:       5000,
		Weight:      2000,
		Status:      true,
		Category_Id: 1,
	}
	domainCategory = productDomain.Category{
		ID:     1,
		Name:   "category_1",
		Status: true,
	}
	os.Exit(m.Run())
}

func TestCheckouId(t *testing.T) {
	t.Run("checkout id", func(t *testing.T) {
		productRepo.On("GetById", mock.AnythingOfType("int")).Return(domainProduct, nil).Once()

		res, err := productService.CheckoutProductId(1)

		assert.NoError(t, err)
		assert.Equal(t, "abcd-efgh-ijkl", res.Code)
	})
}

func TestGetProducts(t *testing.T) {
	t.Run("get all product", func(t *testing.T) {
		productRepo.On("GetProducts").Return([]productDomain.Product{domainProduct}, nil).Once()

		res, err := productService.GetProducts()

		assert.NoError(t, err)
		assert.Equal(t, "abcd-efgh-ijkl", res[0].Code)
	})
}

func TestGetProductId(t *testing.T) {
	t.Run("get product by id", func(t *testing.T) {
		productRepo.On("GetById", mock.AnythingOfType("int")).Return(domainProduct, nil).Once()

		productRepo.On("")

		res, err := productService.GetProductId(1)

		assert.NoError(t, err)
		assert.Equal(t, 1, res.ID)
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("create product", func(t *testing.T) {
		productRepo.On("SaveProduct", mock.AnythingOfType("productDomain.Product")).Return(1, nil).Once()
		productRepo.On("GetById", mock.AnythingOfType("int")).Return(domainProduct, nil).Once()

		res, err := productService.CreateProduct(domainProduct)

		assert.NoError(t, err)
		assert.Equal(t, 1, res.ID)
	})
}

func TestDestroyProduct(t *testing.T) {
	t.Run("create product", func(t *testing.T) {
		productRepo.On("GetById", mock.AnythingOfType("int")).Return(domainProduct, nil).Once()
		productRepo.On("Destroy", mock.AnythingOfType("int")).Return(nil).Once()

		res, err := productService.DestroyProduct(1)

		assert.NoError(t, err)
		assert.Equal(t, 1, res.ID)
	})
}

func TestEditProduct(t *testing.T) {
	t.Run("create product", func(t *testing.T) {
		productRepo.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("productDomain.Product")).Return(nil).Once()
		productRepo.On("GetById", mock.AnythingOfType("int")).Return(domainProduct, nil).Once()

		res, err := productService.EditProduct(1, domainProduct)

		assert.NoError(t, err)
		assert.Equal(t, "product_1", res.Name)
	})
}

func TestGetMinPrice(t *testing.T) {
	t.Run("get min price", func(t *testing.T) {
		productRepo.On("GetProducts").Return([]productDomain.Product{domainProduct, {Price: 3500}}, nil).Once()

		res, err := productService.GetMinPrice()
		serviceProduct.SortMin(res)

		assert.NoError(t, err)
		assert.Equal(t, 3500, res[0].Price)
	})
}

func TestGetMaxPrice(t *testing.T) {
	t.Run("get max price", func(t *testing.T) {
		productRepo.On("GetProducts").Return([]productDomain.Product{domainProduct, {Price: 3500}}, nil).Once()

		res, err := productService.GetMaxPrice()
		serviceProduct.SortMAX(res)

		assert.NoError(t, err)
		assert.Equal(t, 5000, res[0].Price)
	})
}

func TestGetCategory(t *testing.T) {
	t.Run("get category", func(t *testing.T) {
		productRepo.On("GetByNameCategory", mock.AnythingOfType("string")).Return([]productDomain.Product{domainProduct}, nil).Once()

		res, err := productService.GetCategory("category_1")

		assert.NoError(t, err)
		assert.Equal(t, 1, res[0].Category_Id)
	})
}

func TestGetCategoryById(t *testing.T) {
	t.Run("get category by id", func(t *testing.T) {
		productRepo.On("GetCategoryById", mock.AnythingOfType("int")).Return(domainCategory, nil).Once()

		res, err := productService.GetCategoryById(1)

		assert.NoError(t, err)
		assert.Equal(t, "category_1", res.Name)
	})
}

func TestCreateCategory(t *testing.T) {
	t.Run("create category", func(t *testing.T) {
		productRepo.On("SaveCategory", mock.Anything).Return(domainCategory.ID, nil).Once()
		productRepo.On("GetCategoryById", mock.Anything).Return(domainCategory, nil).Once()

		res, err := productService.CreateCategory(domainCategory)

		assert.NoError(t, err)
		assert.Equal(t, "category_1", res.Name)
	})
}
