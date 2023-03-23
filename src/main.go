package main

import (
	"auth-service/src/container_binding"
	"auth-service/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.InitRoutes(r)
	container_binding.SetUbBinding()
	r.Run(":9993")
}
