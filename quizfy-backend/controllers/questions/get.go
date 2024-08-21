package question

import (
	"net/http"
	"strconv"

	question_migrations "github.com/YasserRABIE/QUIZFYv2/migrations/questions_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	// Get quiz ID
	quizID := c.Param("quiz_id")
	quizIDInt, err := strconv.Atoi(quizID)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// Get all questions
	questions, err := question_migrations.GetAll(uint(quizIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	r := response.NewSuccess(questions)
	c.JSON(http.StatusOK, r)
}

func GetByID(c *gin.Context) {
	// Get question ID
	questionID := c.Param("question_id")
	questionIDInt, err := strconv.Atoi(questionID)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// Get question by ID
	question, err := question_migrations.GetByID(uint(questionIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	r := response.NewSuccess(question)
	c.JSON(http.StatusOK, r)
}
