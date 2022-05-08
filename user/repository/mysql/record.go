package mysql

import (
	"time"

	UserDomain "github.com/gozzafadillah/user/domain"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       int
	Name     string
	Username string
	Password string
	Role     string
}

func toDomain(rec Users) UserDomain.Users {
	return UserDomain.Users{
		ID:        int(rec.ID),
		Name:      rec.Name,
		Username:  rec.Username,
		Password:  rec.Password,
		Role:      rec.Role,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func fromDomain(rec UserDomain.Users) Users {
	return Users{
		ID:       int(rec.ID),
		Name:     rec.Name,
		Username: rec.Username,
		Password: rec.Password,
		Role:     rec.Role,
	}
}
