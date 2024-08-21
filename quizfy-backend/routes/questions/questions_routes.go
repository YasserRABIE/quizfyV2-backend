package questions

import (
	question "github.com/YasserRABIE/QUIZFYv2/controllers/questions"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/gin-gonic/gin"
)

func QuestionRoutes(api *gin.RouterGroup) {
	question_routes := api.Group("quiz/:quiz_id/questions", auth.AuthMiddleware)
	{
		question_routes.POST("", question.Create)
		question_routes.GET("/all", question.GetAll)

		question_routes.GET("/:question_id", question.GetByID)
		question_routes.PUT("/:question_id", question.Update)
		question_routes.DELETE("/:question_id", question.Delete)
	}
}
