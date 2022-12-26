package api

import (
	"fmt"
	"macus/internal/pkg/config"
	"macus/internal/router"

	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("==================")
	fmt.Println("Go API REST Running on port: " + conf.Server.Port)
	fmt.Println("Microservice Url: " + conf.Microservice.Url)
	fmt.Println("Version: " + conf.Server.Version)
	fmt.Println("==================")
	_ = web.Run(":" + conf.Server.Port)
}
