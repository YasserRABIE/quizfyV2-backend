package user

import (
	user "github.com/YasserRABIE/QUIZFYv2/controllers/account"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	register_routes := api.Group("/register")
	{
		register_routes.POST("/signup", user.Create)
		register_routes.POST("/login", user.Get)
	}

	account_routes := api.Group("/register", auth.AuthMiddleware)
	{
		account_routes.GET("/account", user.GetValidatedUser)
	}
}
