package dao

import (
	"auth-service/src/models"
)

type UserDao struct{}

func (ud UserDao) FindByEmail(email string) models.User {
	user := models.User{}
	Database.First(&user, "email=?", email)
	return user
}

func (ud UserDao) SaveUser(user models.User) models.User {
	Database.Save(&user)
	return user
}
