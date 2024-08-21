package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// Get token from Authorization header
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		utils.HandleError(c, fmt.Errorf("authorization header missing"), http.StatusUnauthorized)
		return
	}

	// Split the token to check its validity
	tokenParts := strings.Split(bearerToken, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" || tokenParts[1] == "" {
		utils.HandleError(c, fmt.Errorf("invalid token format"), http.StatusUnauthorized)
		return
	}

	token := tokenParts[1]

	// Validate token
	user, err := ValidateToken(token)
	if err != nil {
		utils.HandleError(c, fmt.Errorf("invalid token: %w", err), http.StatusUnauthorized)
		return
	}

	// Set user information in context
	c.Set("user_id", user.ID)
	c.Set("user", user)
	c.Next()
}
