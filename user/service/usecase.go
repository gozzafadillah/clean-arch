package service

import (
	"errors"

	"github.com/gozzafadillah/app/middlewares"
	UserDomain "github.com/gozzafadillah/user/domain"
)

type UserService struct {
	Repository UserDomain.Repository
	jwtAuth    *middlewares.ConfigJwt
}

func NewUserService(repo UserDomain.Repository, JWT *middlewares.ConfigJwt) UserDomain.Service {
	return UserService{
		Repository: repo,
		jwtAuth:    JWT,
	}
}

// ============================ Code Here ==============================
// GetId implements UserDomain.Service
func (us UserService) GetId(id int) (data UserDomain.Users, err error) {
	data, err = us.Repository.GetById(id)
	if err != nil {
		return UserDomain.Users{}, err
	}
	return data, nil
}

// BanUser implements UserDomain.Service
func (us UserService) BanUser(username string) (UserDomain.Users, error) {
	var rec UserDomain.Users
	id, err := us.Repository.GetByUsername(username)
	if err != nil {
		return UserDomain.Users{}, errors.New("user not found, please check again")
	}
	data, err := us.Repository.BanUser(id, rec)
	if err != nil {
		return UserDomain.Users{}, errors.New("user cant change status after get id")
	}
	return data, nil
}

// GetUsername implements UserDomain.Service
func (us UserService) GetUsername(username string) (UserDomain.Users, error) {
	id, err := us.Repository.GetByUsername(username)
	if err != nil {
		return UserDomain.Users{}, errors.New("username not found")
	}
	data, err := us.GetId(id)
	if err != nil {
		return UserDomain.Users{}, errors.New("id username not found")
	}
	return data, nil
}

// CreateToken implements UserDomain.Service
func (us UserService) Login(username string, password string) (string, error) {
	dataUser, err := us.Repository.GetUsernamePassword(username, password)

	if err != nil {
		return "token failed generate", errors.New("username and password missmatch")
	}

	token, err := us.jwtAuth.GenerateToken(dataUser.ID, dataUser.Status)

	if err != nil {
		return "kosong da", errors.New("ada yang salah dengan anda")
	}
	return token, nil
}

// InsertData implements UserDomain.Service
func (us UserService) InsertData(domain UserDomain.Users) (response UserDomain.Users, err error) {
	id, errorResp := us.Repository.Save(domain)

	if errorResp != nil {
		return UserDomain.Users{}, errors.New("can't insert to database")
	}
	record, errorResp2 := us.Repository.GetById(id)
	if errorResp2 != nil {
		return UserDomain.Users{}, errors.New("data not found")
	}
	return record, nil
}

// ======================================================================
