package mysql

import (
	"errors"

	"github.com/gozzafadillah/helper/encryption"
	userDomain "github.com/gozzafadillah/user/domain"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) userDomain.Repository {
	return userRepo{
		DB: db,
	}
}

// =============================== code here ===================================
// GetByUsername implements userDomain.Repository
// UpdateUser implements userDomain.Repository
func (ur userRepo) BanUser(id int, domain userDomain.Users) (userDomain.Users, error) {
	var rec Users
	domain.Status = false
	err := ur.DB.Model(&rec).Where("id = ?", id).Update("status", domain.Status).Error
	ur.DB.Where("id = ?", id).First(&rec)
	return toDomain(rec), err
}

func (ur userRepo) GetByUsername(username string) (int, error) {
	var rec Users
	err := ur.DB.Where("username = ?", username).First(&rec).Error
	return rec.ID, err
}

// GetById implements userDomain.Repository
func (ur userRepo) GetById(id int) (domain userDomain.Users, err error) {
	var newRecord Users
	err = ur.DB.Where("id = ?", id).First(&newRecord).Error
	return toDomain(newRecord), err
}

// GetUsernamePassword implements userDomain.Repository
func (ur userRepo) GetUsernamePassword(username string, password string) (domain userDomain.Users, err error) {
	var record Users
	errResp := ur.DB.Where("username = ?", username).First(&record).Error
	if err := encryption.CheckPasswordHash(password, record.Password); !err && !record.Status {
		return userDomain.Users{}, errors.New("username and password wrong")
	}
	return toDomain(record), errResp
}

// Save implements userDomain.Repository
func (ur userRepo) Save(domain userDomain.Users) (id int, err error) {
	domain.Role = "customer"
	domain.Status = true
	err = ur.DB.Save(&domain).Error
	return domain.ID, err
}

// ===============================================================================
