package handlers

import (
	"macus/pkg/models"
	"macus/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenHandler struct {
	TokenService services.TokenService
}

func (h AuthenHandler) LogIn(c *gin.Context) {
	var login models.LogIn
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		c.Abort()
		return
	}
	accessToken, err := h.TokenService.NewToken(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, accessToken)

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
	// 	Id:        login.UserName,
	// 	ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	// })

	// if tokenSIgn, err := token.SignedString([]byte("MySignatureYSjoSWAQSF")); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"token": tokenSIgn,
	// 	})
	// }
}
