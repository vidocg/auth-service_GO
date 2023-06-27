package controller

import (
	"auth-service/internal/custom_error"
	"auth-service/internal/custom_validator"
	"auth-service/internal/models"
	"auth-service/internal/service"
)

type AuthControllerImpl struct {
	authService service.AuthService
	validator   custom_validator.CustomValidator
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthControllerImpl{authService: authService, validator: custom_validator.NewValidator()}
}

func (ac AuthControllerImpl) GenerateToken(req *models.AuthRequest) (*models.AuthResponse, *custom_error.AppError) {
	return ac.authService.GenerateToken(req)
}

func (ac AuthControllerImpl) SaveUser(req models.UserCreateDto) (*models.UserDto, *custom_error.AppError) {
	err := ac.validator.Validate(&req)
	if err != nil {
		return nil, err
	}
	return ac.authService.SaveUser(req)
}

func (ac AuthControllerImpl) GetUserByToken(tokenString string) (*models.UserDto, *custom_error.AppError) {
	return ac.authService.GetUserByToken(tokenString)
}
