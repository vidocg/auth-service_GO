package service

import (
	"auth-service/src/custom_error"
	"auth-service/src/dao"
	"auth-service/src/models"
	"auth-service/src/util"
	"fmt"
	"github.com/devfeel/mapper"
)

type AuthServiceImpl struct {
	db     dao.UserDatabase
	mapper mapper.IMapper
}

func NewAuthService(db dao.UserDatabase) AuthService {
	return AuthServiceImpl{db, mapper.NewMapper()}
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

func (as AuthServiceImpl) SaveUser(userCreateDto models.UserCreateDto) (*models.UserDto, *custom_error.AppError) {
	hash, err := util.HashPassword(userCreateDto.Password)
	if err != nil {
		return nil, &custom_error.AppError{
			Error:         err,
			Message:       "hashing password error",
			HttpErrorCode: 400,
		}
	}
	userCreateDto.Password = hash
	userToSave := models.User{}

	_ = as.mapper.Mapper(&userCreateDto, &userToSave)

	savedUser := as.db.SaveUser(userToSave)
	dto := &models.UserDto{}
	_ = as.mapper.Mapper(&savedUser, dto)
	return dto, nil
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

func (as AuthServiceImpl) GetUserByToken(tokenString string) (*models.UserDto, *custom_error.AppError) {
	email, err := util.VerifyJwt(tokenString)
	if err != nil {
		return nil, &custom_error.AppError{
			Error:         err,
			Message:       "Jwt is invalid",
			HttpErrorCode: 403,
		}
	}

	user := as.db.FindByEmail(email)
	dto := &models.UserDto{}
	_ = as.mapper.Mapper(&user, dto)
	return dto, nil
}
