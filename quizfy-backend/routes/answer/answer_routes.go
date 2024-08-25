package answer

import (
	"github.com/YasserRABIE/QUIZFYv2/controllers/answer"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/gin-gonic/gin"
)

func AnswerRoutes(api *gin.RouterGroup) {
	answer_routes := api.Group("quiz/:quiz_id/session/:session_id/answers", auth.AuthMiddleware)
	{
		answer_routes.POST("", answer.Create)
		// answer_routes.GET("/all", answer.GetAll)

		// answer_routes.GET("/:answer_id", answer.GetByID)
		// answer_routes.PUT("/:answer_id", answer.Update)
		// answer_routes.DELETE("/:answer_id", answer.Delete)
	}
}
