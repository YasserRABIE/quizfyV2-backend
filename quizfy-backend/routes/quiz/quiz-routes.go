package quiz

import (
	"github.com/YasserRABIE/QUIZFYv2/controllers/quiz"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/gin-gonic/gin"
)

func QuizRoutes(api *gin.RouterGroup) {
	// Create quiz routes
	quiz_routes := api.Group("/quiz", auth.AuthMiddleware)
	{
		quiz_routes.POST("", quiz.Create)
		quiz_routes.GET("/all", quiz.GetAll)

		// ! Not implemented yet
		// quiz_routes.GET("/:id", quiz.GetByID)
		// quiz_routes.PUT("/:id", quiz.Update)
		// quiz_routes.DELETE("/:id", quiz.Delete)
	}
}
