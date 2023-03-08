package service

import (
	"auth-service/src/dao"
	"auth-service/src/models"
	"auth-service/src/util"
	"log"
)

func GenerateToken(req *models.AuthRequest) models.AuthResponse {
	user := getUserByEmail(req.Email)
	if util.CheckPasswordHash(req.Password, user.Password) != true {
		//todo handle error
	}

	token, refreshToken := util.GenerateJwt(user)
	user.RefreshToken = refreshToken
	dao.SaveUser(user)

	return models.AuthResponse{Jwt: token, Refresh: refreshToken}
}

func SaveUser(user models.User) models.User {
	hash, error := util.HashPassword(user.Password)
	if error != nil {
		//todo handle error
		log.Fatal(error)
	}
	user.Password = hash
	return dao.SaveUser(user)
}

func getUserByEmail(email string) models.User {
	return dao.FindByEmail(email)
}

func GetUserByToken(tokenString string) models.User {
	email, err := util.VerifyJwt(tokenString)
	//todo handle error
	if err != nil {
		log.Fatal(err)
	}
	user := dao.FindByEmail(email)
	return user
}
