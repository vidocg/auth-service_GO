package main

import (
	"auth-service/internal/application_context"
	"auth-service/internal/config"
	"auth-service/internal/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @title           Auth service
// @version         1.0
// @description     Microservice that is developed for authorization and authentication purposes
// @host      localhost:9993
func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "."
	}
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		panic("Cannot load config")
	}

	application_context.LoadContext(&conf)

	r := gin.Default()
	http.InitRoutes(r)
	r.Run(":9993")
}
