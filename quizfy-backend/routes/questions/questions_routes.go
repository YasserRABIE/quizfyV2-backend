package questions

import (
	question "github.com/YasserRABIE/QUIZFYv2/controllers/questions"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/gin-gonic/gin"
)

func QuestionRoutes(api *gin.RouterGroup) {
	question_routes := api.Group("/question", auth.AuthMiddleware)
	{
		question_routes.POST("", question.Create)
		// question_routes.GET("/all", question.GetAll)

		// question_routes.GET("/:id", question.GetByID)
		// question_routes.PUT("/:id", question.Update)
		// question_routes.DELETE("/:id", question.Delete)
	}
}
