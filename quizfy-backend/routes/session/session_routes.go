package session

import (
	"github.com/YasserRABIE/QUIZFYv2/controllers/session"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/gin-gonic/gin"
)

func SessionRoutes(api *gin.RouterGroup) {
	session_routes := api.Group("quiz/:quiz_id/session", auth.AuthMiddleware)
	{
		session_routes.POST("", session.Create)
		// session_routes.GET("/all", session.GetAll)

		// session_routes.GET("/:session_id", session.GetByID)
		// session_routes.PUT("/:session_id", session.Update)
		// session_routes.DELETE("/:session_id", session.Delete)
	}
}
