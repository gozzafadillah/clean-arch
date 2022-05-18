package userDomain_test

import (
	"os"
	"testing"

	"github.com/gozzafadillah/app/middlewares"
	userDomain "github.com/gozzafadillah/user/domain"
	userMock "github.com/gozzafadillah/user/domain/mocks"
	serviceUser "github.com/gozzafadillah/user/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userService userDomain.Service
	domain      userDomain.Users
	userRepo    userMock.Repository
)

func TestMain(m *testing.M) {
	userService = serviceUser.NewUserService(&userRepo, &middlewares.ConfigJwt{})
	domain = userDomain.Users{
		ID:       1,
		Name:     "Aziz",
		Username: "Aziz",
		Password: "12345",
		Role:     "customer",
		Status:   true,
	}
	os.Exit(m.Run())
}

func TestGetId(t *testing.T) {
	t.Run("get by id", func(t *testing.T) {
		userRepo.On("GetById", mock.AnythingOfType("int")).Return(domain, nil).Once()

		res, err := userService.GetId(1)

		assert.NoError(t, err)
		assert.Equal(t, "Aziz", res.Name)
	})
}

func TestInsertData(t *testing.T) {
	t.Run("insert data", func(t *testing.T) {
		userRepo.On("Save", mock.AnythingOfType("userDomain.Users")).Return(1, nil).Once()
		userRepo.On("GetById", mock.AnythingOfType("int")).Return(domain, nil).Once()

		res, err := userService.InsertData(domain)

		assert.NoError(t, err)
		assert.Equal(t, "Aziz", res.Name)
	})
}

func TestGetUsername(t *testing.T) {
	t.Run("get by username", func(t *testing.T) {
		userRepo.On("GetByUsername", mock.AnythingOfType("string")).Return(1, nil).Once()
		userRepo.On("GetById", mock.AnythingOfType("int")).Return(domain, nil).Once()
		res, err := userService.GetUsername("Aziz")

		assert.NoError(t, err)
		assert.Equal(t, "Aziz", res.Name)
	})
}

func TestLogin(t *testing.T) {
	t.Run("login", func(t *testing.T) {
		userRepo.On("GetUsernamePassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain, nil).Once()
		userRepo.On("GenerateToken", mock.AnythingOfType("int"), mock.AnythingOfType("bool")).Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwic3RhdHVzIjp0cnVlfQ.vcnqONfIFjRBD5O4a8LDZXNt2afy4rV2NjmpNBDiAqE", nil).Once()
		res, err := userService.Login("Aziz", "12345")

		assert.NoError(t, err)
		assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwic3RhdHVzIjp0cnVlfQ.vcnqONfIFjRBD5O4a8LDZXNt2afy4rV2NjmpNBDiAqE", res)
	})
}

func TestBanUser(t *testing.T) {
	t.Run("login", func(t *testing.T) {
		userRepo.On("GetByUsername", mock.AnythingOfType("string")).Return(1, nil).Once()
		userRepo.On("BanUser", mock.AnythingOfType("int"), mock.AnythingOfType("userDomain.Users")).Return(userDomain.Users{}, nil).Once()
		res, err := userService.BanUser("Aziz")

		assert.NoError(t, err)
		assert.Equal(t, false, res.Status)
	})
}
