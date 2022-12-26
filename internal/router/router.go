package router

import (
	"fmt"
	"io"
	"macus/docs"
	"macus/internal/api/controllers"
	"macus/internal/api/middlewares"
	"macus/internal/pkg/config"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/h2non/gentleman.v2"
)

func Setup() *gin.Engine {
	app := gin.New()
	conf := config.GetConfig()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())
	// app.Use(middlewares.AuthenHeader())

	cli := gentleman.New()
	cli.URL(conf.Microservice.Url)

	// app.Use(middlewares.ForwardHeader(cli))
	todoHandler := controllers.NewTodoHandler(cli)
	todoV1 := app.Group("/api/v1", middlewares.AuthenHeader(), middlewares.ForwardHeader(cli))
	{
		todo := todoV1.Group("/todos")
		{
			todo.GET("", todoHandler.GetAll)
		}
	}

	app.GET("/healthcheck", controllers.HealthCheck)

	docs.SwaggerInfo.BasePath = "/api/v1"
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
