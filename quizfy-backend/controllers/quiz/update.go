package quiz

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/YasserRABIE/QUIZFYv2/migrations/quiz_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var quizData quiz.Quiz

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// Bind request body
	err = c.ShouldBindJSON(&quizData)
	if err != nil {
		err = errors.New("يرجى ملأ كل الحقول")
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// Update quiz in database
	err = quiz_migrations.Update(uint(idInt), &quizData)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// Send success response
	r := response.NewSuccess("تم تحديث الاختبار بنجاح")
	c.JSON(http.StatusOK, r)
}
