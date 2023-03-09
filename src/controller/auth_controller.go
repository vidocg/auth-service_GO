package controller

import (
	"auth-service/src/models"
	"auth-service/src/service"
)

func GenerateToken(req *models.AuthRequest) models.AuthResponse {
	return service.GenerateToken(req)
}

func SaveUser(req models.User) models.User {
	return service.SaveUser(req)
}

func GetUserByToken(tokenString string) models.User {
	return service.GetUserByToken(tokenString)
}
