package controller

import (
	"auth-service/src/error"
	"auth-service/src/models"
	"auth-service/src/service"
)

func GenerateToken(req *models.AuthRequest) (*models.AuthResponse, *error.AppError) {
	return service.GenerateToken(req)
}

func SaveUser(req models.User) (*models.User, *error.AppError) {
	return service.SaveUser(req)
}

func GetUserByToken(tokenString string) (*models.User, *error.AppError) {
	return service.GetUserByToken(tokenString)
}
