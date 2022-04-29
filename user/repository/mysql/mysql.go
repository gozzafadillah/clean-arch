package mysql

import (
	UserDomain "github.com/gozzafadillah/user/domain"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

// =============================== code here ===================================

// GetById implements UserDomain.Repository
func (ur userRepo) GetById(id int) (domain UserDomain.Users, err error) {
	newRecord := Users{}
	err = ur.DB.Find("id = ?", id).First(&newRecord).Error

	return toDomain(newRecord), err
}

// GetUsernamePassword implements UserDomain.Repository
func (ur userRepo) GetUsernamePassword(username string, password string) (domain UserDomain.Users, err error) {
	panic("unimplemented")
}

// Save implements UserDomain.Repository
func (ur userRepo) Save(domain UserDomain.Users) (id int, err error) {
	var newRecord Users
	err = ur.DB.Find("id = ?", id).First(&newRecord).Error
	return newRecord.ID, err
}

// ===============================================================================

func NewUserRepository(db *gorm.DB) UserDomain.Repository {
	return userRepo{
		DB: db,
	}
}
