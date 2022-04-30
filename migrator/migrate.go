package migrate

import (
	userRepo "github.com/gozzafadillah/user/repository/mysql"
	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {

	DB.AutoMigrate(&userRepo.Users{})
}
