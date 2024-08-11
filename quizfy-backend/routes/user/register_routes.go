package routes

import (
	"github.com/YasserRABIE/QUIZFYv2/controllers/quizzer"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	register_routes := api.Group("/register")
	{
		// quizzer
		register_routes.POST("/signup/quizzer", quizzer.Create)
		register_routes.POST("/login/quizzer", quizzer.Create)

		// student
		register_routes.POST("/signup/student", quizzer.Create)
		register_routes.POST("/login/student", quizzer.Create)
	}
}
