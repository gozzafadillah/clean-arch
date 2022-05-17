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
func (us UserService) GetId(id int) (response UserDomain.Users, err error) {
	response, err = us.Repository.GetById(id)
	if err != nil {
		return UserDomain.Users{}, err
	}
	return response, nil
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
