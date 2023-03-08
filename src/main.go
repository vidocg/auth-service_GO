package main

import (
	"auth-service/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.InitRoutes(r)
	r.Run(":9993")
}
