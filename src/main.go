package main

import (
	"github.com/gin-gonic/gin"
	"auth-service/src/controller"
)

func main() {
	r := gin.Default()
	controller.InitRoutes(r)
	r.Run(":9993")
}
