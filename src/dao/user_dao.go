package dao

import "auth-service/src/models"

type UserDao struct{}

func (ud UserDao) FindByEmail(email string) models.User {
	return commits[email]
}

func (ud UserDao) SaveUser(user models.User) models.User {
	commits[user.Email] = user
	return user
}

var commits = map[string]models.User{
	"email1": {
		Password: "pass1",
		Email:    "email1",
	},
	"email2": {
		Password: "email2",
		Email:    "email2",
	},
	"email3": {
		Password: "email3",
		Email:    "email3",
	},
	"email4": {
		Password: "email4",
		Email:    "email4",
	},
}
