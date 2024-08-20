package quiz

import (
	"net/http"
	"strconv"

	"github.com/YasserRABIE/QUIZFYv2/migrations/quiz_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

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

func GetByID(c *gin.Context) {
	id := c.Param("id")
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
