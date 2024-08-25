package result

import (
	"github.com/YasserRABIE/QUIZFYv2/controllers/result"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/gin-gonic/gin"
)

func ResultRoutes(api *gin.RouterGroup) {
	result_routes := api.Group("quiz/:quiz_id/session/:session_id/result", auth.AuthMiddleware)
	{
		result_routes.POST("", result.Create)
		result_routes.GET("", result.Get)
	}
}
