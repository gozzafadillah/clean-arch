package migrate

import (
	"github.com/gozzafadillah/app/config"
	userRepo "github.com/gozzafadillah/user/repository/mysql"
)

func AutoMigrate() {
	config.DB.AutoMigrate(&userRepo.Users{})
}
