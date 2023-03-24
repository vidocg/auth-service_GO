package service

import (
	"auth-service/src/custom_error"
	"auth-service/src/dao"
	"auth-service/src/models"
	"auth-service/src/util"
	"fmt"
)

type AuthServiceImpl struct {
	db dao.UserDatabase
}

func NewAuthService(db dao.UserDatabase) AuthService {
	return AuthServiceImpl{db}
}

func (as AuthServiceImpl) GenerateToken(req *models.AuthRequest) (*models.AuthResponse, *custom_error.AppError) {
	user, appErr := as.getUserByEmail(req.Email)

	if appErr != nil {
		return nil, appErr
	}
	if util.CheckPasswordHash(req.Password, user.Password) != true {
		return nil, &custom_error.AppError{
			Error:         fmt.Errorf("wrong password"),
			Message:       "wrong password",
			HttpErrorCode: 401,
		}
	}

	token, refreshToken := util.GenerateJwt(*user)
	user.RefreshToken = refreshToken
	as.db.SaveUser(*user)

	return &models.AuthResponse{Jwt: token, Refresh: refreshToken}, nil
}

func (as AuthServiceImpl) SaveUser(user models.User) (*models.User, *custom_error.AppError) {
	hash, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, &custom_error.AppError{
			Error:         err,
			Message:       "hashing password error",
			HttpErrorCode: 400,
		}
	}
	user.Password = hash
	savedUser := as.db.SaveUser(user)
	return &savedUser, nil
}

func (as AuthServiceImpl) getUserByEmail(email string) (*models.User, *custom_error.AppError) {
	user := as.db.FindByEmail(email)
	if &user == nil {
		return nil, &custom_error.AppError{
			Error:         fmt.Errorf("user not found"),
			Message:       "user not found",
			HttpErrorCode: 404,
		}
	}
	return &user, nil
}

func (as AuthServiceImpl) GetUserByToken(tokenString string) (*models.User, *custom_error.AppError) {
	email, err := util.VerifyJwt(tokenString)
	if err != nil {
		return nil, &custom_error.AppError{
			Error:         err,
			Message:       "Jwt is invalid",
			HttpErrorCode: 403,
		}
	}

	user := as.db.FindByEmail(email)
	return &user, nil
}
