package dao

import "auth-service/internal/models"

type UserDatabase interface {
	FindByEmail(email string) models.User
	SaveUser(user models.User) models.User
}
