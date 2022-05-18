package serviceUser

import (
	"errors"

	"github.com/gozzafadillah/app/middlewares"
	userDomain "github.com/gozzafadillah/user/domain"
)

type UserService struct {
	Repository userDomain.Repository
	jwtAuth    *middlewares.ConfigJwt
}

func NewUserService(repo userDomain.Repository, JWT *middlewares.ConfigJwt) userDomain.Service {
	return UserService{
		Repository: repo,
		jwtAuth:    JWT,
	}
}

// ============================ Code Here ==============================
// GetId implements userDomain.Service
func (us UserService) GetId(id int) (data userDomain.Users, err error) {
	data, err = us.Repository.GetById(id)
	if err != nil {
		return userDomain.Users{}, err
	}
	return data, nil
}

// BanUser implements userDomain.Service
func (us UserService) BanUser(username string) (userDomain.Users, error) {
	var rec userDomain.Users
	id, err := us.Repository.GetByUsername(username)
	if err != nil {
		return userDomain.Users{}, errors.New("user not found, please check again")
	}
	data, err := us.Repository.BanUser(id, rec)
	if err != nil {
		return userDomain.Users{}, errors.New("user cant change status after get id")
	}
	return data, nil
}

// GetUsername implements userDomain.Service
func (us UserService) GetUsername(username string) (userDomain.Users, error) {
	id, err := us.Repository.GetByUsername(username)
	if err != nil {
		return userDomain.Users{}, errors.New("username not found")
	}
	data, err := us.GetId(id)
	if err != nil {
		return userDomain.Users{}, errors.New("id username not found")
	}
	return data, nil
}

// CreateToken implements userDomain.Service
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

// InsertData implements userDomain.Service
func (us UserService) InsertData(domain userDomain.Users) (response userDomain.Users, err error) {
	id, err := us.Repository.Save(domain)

	if err != nil {
		return userDomain.Users{}, errors.New("can't insert to database")
	}
	record, err := us.Repository.GetById(id)
	if err != nil {
		return userDomain.Users{}, errors.New("data not found")
	}
	return record, nil
}

// ======================================================================
