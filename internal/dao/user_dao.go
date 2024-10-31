package dao

import (
	"auth-service/internal/custom_error"
	"auth-service/internal/models"
	"fmt"

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

func (ud UserDao) SaveUser(user models.User) (models.User, *custom_error.AppError) {
	error := ud.db.Save(&user).Error
	if error != nil {
		return user, &custom_error.AppError{
			Error:         fmt.Errorf("cannot save user"),
			Message:       "User is not valid",
			HttpErrorCode: 400,
		}
	}
	return user, nil
}
