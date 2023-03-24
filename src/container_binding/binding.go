package container_binding

import (
	"auth-service/src/controller"
	"auth-service/src/dao"
	"auth-service/src/service"
	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

func SetUbBinding(db *gorm.DB) {
	userDao := dao.NewUserDao(db)
	authService := service.NewAuthService(userDao)
	authController := controller.NewAuthController(authService)

	container.Singleton(func() controller.AuthController {
		return authController
	})
}
