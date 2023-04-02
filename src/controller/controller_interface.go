package controller

import (
	"auth-service/src/custom_error"
	"auth-service/src/models"
)

type AuthController interface {
	GenerateToken(req *models.AuthRequest) (*models.AuthResponse, *custom_error.AppError)
	SaveUser(req models.User) (*models.User, *custom_error.AppError)
	GetUserByToken(tokenString string) (*models.UserDto, *custom_error.AppError)
}
