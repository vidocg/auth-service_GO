package main

import (
	"auth-service/src/application_context"
	"auth-service/src/config"
	"github.com/gin-gonic/gin"
)

// @title           Auth service
// @version         1.0
// @description     Microservice that is developed for authorization and authentication purposes

// @host      localhost:9993
func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		panic("Cannot load config")
	}

	application_context.LoadContext(&conf)

	r := gin.Default()
	InitRoutes(r)
	r.Run(":9993")
}
