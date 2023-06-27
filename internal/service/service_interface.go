package service

import (
	"auth-service/internal/custom_error"
	"auth-service/internal/models"
)

type AuthService interface {
	GenerateToken(req *models.AuthRequest) (*models.AuthResponse, *custom_error.AppError)
	SaveUser(user models.UserCreateDto) (*models.UserDto, *custom_error.AppError)
	GetUserByToken(tokenString string) (*models.UserDto, *custom_error.AppError)
	LogInThroughSocialNetwork(user models.SocialNetworkUser) *models.AuthResponse
}
