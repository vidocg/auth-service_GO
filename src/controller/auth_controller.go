package controller

import (
	"auth-service/src/custom_error"
	"auth-service/src/models"
	"auth-service/src/service"
)

type AuthControllerImpl struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthControllerImpl{authService}
}

func (ac AuthControllerImpl) GenerateToken(req *models.AuthRequest) (*models.AuthResponse, *custom_error.AppError) {
	return ac.authService.GenerateToken(req)
}

func (ac AuthControllerImpl) SaveUser(req models.User) (*models.User, *custom_error.AppError) {
	return ac.authService.SaveUser(req)
}

func (ac AuthControllerImpl) GetUserByToken(tokenString string) (*models.UserDto, *custom_error.AppError) {
	return ac.authService.GetUserByToken(tokenString)
}
