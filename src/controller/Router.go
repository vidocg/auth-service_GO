package controller

import (
	"auth-service/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/token", func(context *gin.Context) {
		authRequest:= &models.AuthRequest{}
		err := context.BindJSON(authRequest)
		if err != nil  {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		context.JSON(http.StatusOK, GenerateToken(authRequest))
	})
}
