package mysql

import (
	"time"

	userDomain "github.com/gozzafadillah/user/domain"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       int
	Name     string
	Username string
	Password string
	Role     string
	Status   bool
}

func toDomain(rec Users) userDomain.Users {
	return userDomain.Users{
		ID:        int(rec.ID),
		Name:      rec.Name,
		Username:  rec.Username,
		Password:  rec.Password,
		Role:      rec.Role,
		Status:    rec.Status,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
