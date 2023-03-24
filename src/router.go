package main

import (
	"auth-service/src/application_context"
	"auth-service/src/custom_error"
	"auth-service/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(r *gin.Engine) {
	controller := application_context.ResolveAuthController()
	r.POST("/token", func(context *gin.Context) {
		authRequest := &models.AuthRequest{}
		err := context.BindJSON(authRequest)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		obj, saveErr := controller.GenerateToken(authRequest)
		resolveResponse(obj, saveErr, context)
	})

	r.POST("/user", func(context *gin.Context) {
		user := &models.User{}
		err := context.BindJSON(user)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		obj, saveErr := controller.SaveUser(*user)
		resolveResponse(obj, saveErr, context)
	})

	r.GET("/user", func(context *gin.Context) {
		token := context.Query("token")
		if &token == nil {
			context.AbortWithError(http.StatusBadRequest, fmt.Errorf("jwt is null"))
			return
		}
		obj, err := controller.GetUserByToken(token)
		resolveResponse(obj, err, context)
	})
}

func resolveResponse(obj any, err *custom_error.AppError, context *gin.Context) {
	if err != nil {
		context.JSON(err.HttpErrorCode, err.Message)
	} else {
		context.JSON(http.StatusOK, obj)
	}
}
