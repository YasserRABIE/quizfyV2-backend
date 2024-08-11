package routes

import (
	"github.com/YasserRABIE/QUIZFYv2/controllers/register/quizzer"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/api/quizzer", quizzer.Create)
}
