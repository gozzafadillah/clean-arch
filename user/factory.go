package user

import (
	"gorm.io/gorm"

	userAPI "github.com/gozzafadillah/user/handler/api"
	userRepoMysql "github.com/gozzafadillah/user/repository/mysql"
	userService "github.com/gozzafadillah/user/service"
)

func NewUserFactory(db *gorm.DB) (userHandler userAPI.UserHandler) {
	userRepo := userRepoMysql.NewUserRepository(db)
	userServ := userService.NewUserService(userRepo)
	userHandler = userAPI.NewUserHandler(userServ)
	return
}
