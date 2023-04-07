package application_context

import (
	"auth-service/src/config"
	"auth-service/src/controller"
	"auth-service/src/dao"
	"auth-service/src/service"
	"github.com/golobby/container/v3"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func LoadContext(config *config.Config) {
	db := dao.Configure(config)
	userDao := dao.NewUserDao(db)
	authService := service.NewAuthService(userDao)
	authController := controller.NewAuthController(authService)

	goth.UseProviders(
		google.New(config.GoogleClientId, config.GoogleClientSecret, config.GoogleAuthCallbackUrl, "email", "profile"),
	)

	container.Singleton(func() controller.AuthController {
		return authController
	})
}
