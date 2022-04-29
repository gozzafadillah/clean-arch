package mysql

import (
	UserDomain "github.com/gozzafadillah/user/domain"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       int
	Name     string
	Username string
	Password string
}

func toDomain(rec Users) UserDomain.Users {
	return UserDomain.Users{
		ID:       int(rec.ID),
		Name:     rec.Name,
		Username: rec.Username,
		Password: rec.Password,
	}
}

func fromDomain(rec UserDomain.Users) Users {
	return Users{
		ID:       int(rec.ID),
		Name:     rec.Name,
		Username: rec.Username,
		Password: rec.Password,
	}
}
