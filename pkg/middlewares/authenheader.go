package middlewares

import (
	"fmt"
	"macus/pkg/models"
	"macus/pkg/models/apperrors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// used to help extract validation errors
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func AuthenHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header models.RequireHeader
		if err := c.ShouldBindHeader(&header); err != nil {
			fmt.Println("AuthenHeader=>" + header.XAuth)
			fmt.Println("AuthenHeader=>" + header.XToken)
			if errs, ok := err.(validator.ValidationErrors); ok {
				var invalidArgs []invalidArgument

				for _, err := range errs {
					invalidArgs = append(invalidArgs, invalidArgument{
						err.Field(),
						err.Value().(string),
						err.Tag(),
						err.Param(),
					})
				}

				err := apperrors.NewBadRequest("Invalid request parameters. See invalidArgs")

				c.JSON(err.Status(), gin.H{
					"error":       err,
					"invalidArgs": invalidArgs,
				})
				c.Abort()
				return
			}

			// otherwise error type is unknown
			err := apperrors.NewAuthorization("header is required")
			c.JSON(err.Status(), gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		fmt.Println("AuthenHeader=>" + header.XAuth)
		fmt.Println("AuthenHeader=>" + header.XToken)
		c.Next()
	}
}
