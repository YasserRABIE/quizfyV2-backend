package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")

	tokenParts := strings.Split(bearerToken, " ")
	tokenErr := len(tokenParts) == 1 || tokenParts[0] != "Bearer" || tokenParts[1] == ""
	if tokenErr {
		utils.HandleError(c, fmt.Errorf("invalid token"), http.StatusUnauthorized)
		return
	}

	token := tokenParts[1]
	user, err := ValidateToken(token)
	utils.HandleError(c, err, http.StatusUnauthorized)

	c.Set("user", user)
	c.Next()
}
