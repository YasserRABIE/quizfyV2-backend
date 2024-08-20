package question

import (
	"net/http"

	question_migrations "github.com/YasserRABIE/QUIZFYv2/migrations/questions_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var question quiz.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// Set the QuestionID for each option
	for i := range question.Options {
		question.Options[i].QuestionID = question.ID
	}

	if err := question_migrations.Create(&question); err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, question)
}
