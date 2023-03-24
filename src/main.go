package main

import (
	"auth-service/src/config"
	"auth-service/src/container_binding"
	"auth-service/src/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	conf, err := config.LoadConfig(".")
	if err != nil {
		panic("Cannot load config")
	}

	db := dao.Configure(&conf)

	container_binding.SetUbBinding(db)

	InitRoutes(r)
	r.Run(":9993")
}
