package main

import (
	"auth-service/src/config"
	"auth-service/src/container_binding"
	"auth-service/src/controller"
	"auth-service/src/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.InitRoutes(r)
	container_binding.SetUbBinding()
	conf, err := config.LoadConfig(".")
	if err != nil {
		panic("Cannot load config")
	}
	dao.Configure(&conf)
	r.Run(":9993")
}
