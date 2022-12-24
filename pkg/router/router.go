package router

import (
	"macus/configs"
	"macus/pkg/handlers"
	"macus/pkg/middlewares"
	"macus/pkg/services"

	"github.com/gin-gonic/gin"
	"gopkg.in/h2non/gentleman.v2"
)

func AddAPIs(engine *gin.Engine) {
	cli := gentleman.New()
	// cli.URL("https://jsonplaceholder.typicode.com")
	cli.URL("http://localhost:5010")
	engine.Use(middlewares.AuthenHeader())
	engine.Use(middlewares.ForwardHeader(cli))
	todoHandler := handlers.NewTodoHandler(cli)
	todoV1 := engine.Group("/api/v1")
	{
		todo := todoV1.Group("/todos")
		{
			todo.GET("", todoHandler.GetAll)
		}
	}

	privKey := configs.LoadPrivateKey()
	pubKey := configs.LoadPublicKey()
	idExp := configs.LoadTokenExp()
	authenHandler := handlers.AuthenHandler{
		TokenService: services.NewTokenService(&services.TSConfig{
			PrivKey:          privKey,
			PubKey:           pubKey,
			IDExpirationSecs: idExp,
		}),
	}

	authenv1 := engine.Group("/api/v1")
	{
		authen := authenv1.Group("/authen")
		{
			authen.POST("/LogIn", authenHandler.LogIn)
		}
	}

	engine.GET("/healthcheck", handlers.HealthCheck)
}
