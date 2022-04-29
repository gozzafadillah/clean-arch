package service

import (
	errorConv "github.com/gozzafadillah/helper/error"
	UserDomain "github.com/gozzafadillah/user/domain"
)

type userService struct {
	Repository UserDomain.Repository
}

// ============================ Code Here ==============================

// CreateToken implements UserDomain.Service
func (userService) CreateToken(username string, password string) string {
	panic("unimplemented")
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

func NewUserService(repo UserDomain.Repository) UserDomain.Service {
	return userService{
		Repository: repo,
	}
}
