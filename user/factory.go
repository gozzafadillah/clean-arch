package user

import (
	"gorm.io/gorm"

	_middleware "github.com/gozzafadillah/app/middlewares"

	userAPI "github.com/gozzafadillah/user/handler/api"
	userRepoMysql "github.com/gozzafadillah/user/repository/mysql"
	userService "github.com/gozzafadillah/user/service"
)

func NewUserFactory(db *gorm.DB) (userHandler userAPI.UserHandler) {

	configJWT := _middleware.ConfigJwt{
		SecretJWT: "2345",
	}

	userRepo := userRepoMysql.NewUserRepository(db)
	userServ := userService.NewUserService(userRepo, &configJWT)
	userHandler = userAPI.NewUserHandler(userServ)
	return
}
