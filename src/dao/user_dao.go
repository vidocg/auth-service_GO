package dao

import (
	"auth-service/src/models"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return UserDao{db}
}

func (ud UserDao) FindByEmail(email string) models.User {
	user := models.User{}
	ud.db.First(&user, "email=?", email)
	return user
}

func (ud UserDao) SaveUser(user models.User) models.User {
	ud.db.Save(&user)
	return user
}
