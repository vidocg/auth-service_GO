package application_context

import (
	"auth-service/src/config"
	"auth-service/src/controller"
	"auth-service/src/dao"
	"auth-service/src/service"
	"github.com/golobby/container/v3"
)

func LoadContext(config *config.Config) {
	db := dao.Configure(config)
	userDao := dao.NewUserDao(db)
	authService := service.NewAuthService(userDao)
	authController := controller.NewAuthController(authService)

	container.Singleton(func() controller.AuthController {
		return authController
	})
}
