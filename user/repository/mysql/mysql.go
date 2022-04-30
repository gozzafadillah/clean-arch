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
	var newRecord Users
	err = ur.DB.Where("id = ?", id).Find(&newRecord).Error
	return toDomain(newRecord), err
}

// GetUsernamePassword implements UserDomain.Repository
func (ur userRepo) GetUsernamePassword(username string, password string) (domain UserDomain.Users, err error) {
	var record Users
	errResp := ur.DB.Where("username = ? AND password = ?", username, password).Find(&record).Error
	return toDomain(record), errResp
}

// Save implements UserDomain.Repository
func (ur userRepo) Save(domain UserDomain.Users) (id int, err error) {
	err = ur.DB.Save(&domain).Error
	return domain.ID, err
}

// ===============================================================================

func NewUserRepository(db *gorm.DB) UserDomain.Repository {
	return userRepo{
		DB: db,
	}
}
