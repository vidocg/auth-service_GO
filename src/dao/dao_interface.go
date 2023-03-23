package dao

import "auth-service/src/models"

type UserDatabase interface {
	FindByEmail(email string) models.User
	SaveUser(user models.User) models.User
}
