package service

import (
	"github.com/gozzafadillah/app/middlewares"
	errorConv "github.com/gozzafadillah/helper/error"
	UserDomain "github.com/gozzafadillah/user/domain"
)

type userService struct {
	Repository UserDomain.Repository
	jwtAuth    *middlewares.ConfigJwt
}

// ============================ Code Here ==============================

// CreateToken implements UserDomain.Service
func (us userService) Login(username string, password string) (string, error) {
	dataUser, err := us.Repository.GetUsernamePassword(username, password)
	if err != nil {
		return "", errorConv.Conversion(err)
	}
	token := us.jwtAuth.GenerateToken(dataUser.ID)
	return token, nil
}

// InsertData implements UserDomain.Service
func (us userService) InsertData(domain UserDomain.Users) (response UserDomain.Users, err error) {
	id, errorResp := us.Repository.Save(domain)

	if errorResp != nil {
		return UserDomain.Users{}, errorConv.Conversion(errorResp)
	}
	record, errorResp2 := us.Repository.GetById(id)
	if errorResp2 != nil {
		return UserDomain.Users{}, errorConv.Conversion(errorResp2)
	}
	return record, nil
}

// ======================================================================

func NewUserService(repo UserDomain.Repository, JWT *middlewares.ConfigJwt) UserDomain.Service {
	return userService{
		Repository: repo,
		jwtAuth:    JWT,
	}
}
