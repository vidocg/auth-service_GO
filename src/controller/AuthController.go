package controller

import (
	"auth-service/src/models"
)

func GenerateToken(req *models.AuthRequest) models.AuthResponse {
	return models.AuthResponse{Jwt: req.Email + req.Password, Refresh: req.Password + req.Email}
}
