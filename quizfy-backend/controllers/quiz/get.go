package quiz

import (
	"net/http"
	"strconv"

	"github.com/YasserRABIE/QUIZFYv2/migrations/quiz_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

// GetAll get all quizzes
func GetAll(c *gin.Context) {
	userID, _ := c.Get("user_id")

	quizzes, err := quiz_migrations.GetAll(userID.(uint))
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	r := response.NewSuccess(quizzes)
	c.JSON(http.StatusOK, r)
}

// GetByID get quiz by id
func GetByID(c *gin.Context) {
	id := c.Param("quiz_id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	quiz, err := quiz_migrations.GetByID(uint(idInt))
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	r := response.NewSuccess(quiz)
	c.JSON(http.StatusOK, r)
}
// ========================================
// quizzer
// ========================================

// GetByQuizzerID get quiz by quizzer id and state
func GetExamsByQuizzerID(c *gin.Context) {
	quizzerID := c.MustGet("user_id").(uint)
	state := c.Query("state")

	quiz, err := quiz_migrations.GetExamsByQuizzerID(quizzerID, state)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	r := response.NewSuccess(quiz)
	c.JSON(http.StatusOK, r)
}
