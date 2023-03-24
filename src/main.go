package main

import (
	"auth-service/src/application_context"
	"auth-service/src/config"
	"github.com/gin-gonic/gin"
)

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
