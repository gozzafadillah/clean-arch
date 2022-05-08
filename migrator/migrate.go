package migrate

import (
	productRepo "github.com/gozzafadillah/product/repository/mysql"
	userRepo "github.com/gozzafadillah/user/repository/mysql"
	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {

	DB.AutoMigrate(&userRepo.Users{}, &productRepo.Products{})
}
