package controller

import (
	"auth-service/internal/custom_error"
	"auth-service/internal/models"
)

type AuthController interface {
	GenerateToken(req *models.AuthRequest) (*models.AuthResponse, *custom_error.AppError)
	SaveUser(req models.UserCreateDto) (*models.UserDto, *custom_error.AppError)
	GetUserByToken(tokenString string) (*models.UserDto, *custom_error.AppError)
}
