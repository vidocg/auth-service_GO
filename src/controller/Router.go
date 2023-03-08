package controller

import (
	"auth-service/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/token", func(context *gin.Context) {
		authRequest := &models.AuthRequest{}
		err := context.BindJSON(authRequest)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, GenerateToken(authRequest))
	})

	r.POST("/user", func(context *gin.Context) {
		user := &models.User{}
		err := context.BindJSON(user)
		if err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, SaveUser(*user))
	})

	r.GET("/user", func(context *gin.Context) {
		token := context.Query("token")
		if token == "" {
			context.AbortWithError(http.StatusBadRequest, fmt.Errorf("jwt is null"))
			return
		}
		context.JSON(http.StatusOK, GetUserByToken(token))
	})
}
