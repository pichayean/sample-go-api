package main

import (
	"log"
	"macus/pkg/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(os.Getenv("APP_ENV") + ".env")
	engine := gin.Default()
	router.AddAPIs(engine)
	log.Fatal((engine.Run(":" + os.Getenv("PORT"))))
}
