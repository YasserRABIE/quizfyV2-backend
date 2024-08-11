package routes

import (
	"github.com/YasserRABIE/QUIZFYv2/controllers/quizzer"
	"github.com/YasserRABIE/QUIZFYv2/controllers/student"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	register_routes := api.Group("/register")
	{
		// quizzer
		register_routes.POST("/signup/quizzer", quizzer.Create)
		register_routes.POST("/login/quizzer", quizzer.Get)

		// student
		register_routes.POST("/signup/student", student.Create)
		register_routes.POST("/login/student", student.Get)
	}
}
