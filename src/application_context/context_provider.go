package application_context

import (
	"auth-service/src/controller"
	"auth-service/src/service"
	"auth-service/src/util"
	"github.com/golobby/container/v3"
)

func ResolveAuthController() controller.AuthController {
	var authController controller.AuthController
	containerErr := container.Resolve(&authController)
	if containerErr != nil {
		panic("AuthController impl is not fount")
	}

	return authController
}

func ResolveAuthService() service.AuthService {
	var authService service.AuthService
	containerErr := container.Resolve(&authService)
	if containerErr != nil {
		panic("AuthService impl is not fount")
	}

	return authService
}

func ResolveLogger() util.CustomLogger {
	var logger util.CustomLogger
	containerErr := container.Resolve(&logger)
	if containerErr != nil {
		panic("Custom logger impl is not fount")
	}

	return logger
}
