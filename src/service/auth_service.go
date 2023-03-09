package service

import (
	"auth-service/src/dao"
	"auth-service/src/error"
	"auth-service/src/models"
	"auth-service/src/util"
	"fmt"
)

func GenerateToken(req *models.AuthRequest) (*models.AuthResponse, *error.AppError) {
	user, err := getUserByEmail(req.Email)

	if err != nil {
		return nil, err
	}
	if util.CheckPasswordHash(req.Password, user.Password) != true {
		return nil, &error.AppError{
			Error:         fmt.Errorf("wrong password"),
			Message:       "wrong password",
			HttpErrorCode: 401,
		}
	}

	token, refreshToken := util.GenerateJwt(*user)
	user.RefreshToken = refreshToken
	dao.SaveUser(*user)

	return &models.AuthResponse{Jwt: token, Refresh: refreshToken}, nil
}

func SaveUser(user models.User) (*models.User, *error.AppError) {
	hash, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, &error.AppError{
			Error:         err,
			Message:       "hashing password error",
			HttpErrorCode: 400,
		}
	}
	user.Password = hash
	savedUser := dao.SaveUser(user)
	return &savedUser, nil
}

func getUserByEmail(email string) (*models.User, *error.AppError) {
	user := dao.FindByEmail(email)
	if &user == nil {
		return nil, &error.AppError{
			Error:         fmt.Errorf("user not found"),
			Message:       "user not found",
			HttpErrorCode: 404,
		}
	}
	return &user, nil
}

func GetUserByToken(tokenString string) (*models.User, *error.AppError) {
	email, err := util.VerifyJwt(tokenString)
	//todo handle error
	if err != nil {
		return nil, &error.AppError{
			Error:         err,
			Message:       "Jwt is invalid",
			HttpErrorCode: 403,
		}
	}
	user := dao.FindByEmail(email)
	return &user, nil
}
