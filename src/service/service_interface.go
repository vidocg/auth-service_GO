package service

import (
	"auth-service/src/custom_error"
	"auth-service/src/models"
)

type AuthService interface {
	GenerateToken(req *models.AuthRequest) (*models.AuthResponse, *custom_error.AppError)
	SaveUser(user models.UserCreateDto) (*models.UserDto, *custom_error.AppError)
	GetUserByToken(tokenString string) (*models.UserDto, *custom_error.AppError)
}
