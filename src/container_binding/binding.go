package container_binding

import (
	"auth-service/src/dao"
	"auth-service/src/error"
	"github.com/golobby/container/v3"
)

func SetUbBinding() {
	container.Singleton(func() dao.UserDatabase {
		return &dao.UserDao{}
	})
}

func ResolveUserDao() (dao.UserDatabase, *error.AppError) {
	var db dao.UserDatabase
	containerErr := container.Resolve(&db)
	if containerErr != nil {
		return nil, &error.AppError{
			Error:         containerErr,
			Message:       "UserDatabase impl is not fount",
			HttpErrorCode: 501,
		}
	}

	return db, nil
}
