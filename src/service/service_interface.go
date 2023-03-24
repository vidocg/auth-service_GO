package service

import (
	"auth-service/src/custom_error"
	"auth-service/src/models"
)

type AuthService interface {
	GenerateToken(req *models.AuthRequest) (*models.AuthResponse, *custom_error.AppError)
	SaveUser(user models.User) (*models.User, *custom_error.AppError)
	GetUserByToken(tokenString string) (*models.User, *custom_error.AppError)
}
