package middlewares

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
)

func ForwardHeader(cli *gentleman.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		cli.Use(headers.Set("XToken", c.Request.Header.Get("XAuth")))
		cli.Use(headers.Set("XAuth", c.Request.Header.Get("XToken")))
	}
}
