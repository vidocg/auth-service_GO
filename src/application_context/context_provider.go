package application_context

import (
	"auth-service/src/controller"
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
