package routes

import (
	user "github.com/YasserRABIE/QUIZFYv2/controllers/account"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	register_routes := api.Group("/register")
	{
		register_routes.POST("/signup", user.Create)
		register_routes.POST("/login", user.Get)
	}
}
