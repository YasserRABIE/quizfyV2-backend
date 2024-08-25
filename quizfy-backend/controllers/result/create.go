package result

import (
	"net/http"
	"strconv"

	"github.com/YasserRABIE/QUIZFYv2/migrations/result_migrations"
	"github.com/YasserRABIE/QUIZFYv2/migrations/session_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"github.com/YasserRABIE/QUIZFYv2/models/response"

	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var result quiz.Result

	sessionID, quizID, err := getIdsAsUint(c)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}
	if err := c.ShouldBindJSON(&result.ReviewedAnswers); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// finish the session
	_, err = session_migrations.Finish(sessionID)
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	err = result_migrations.Create(sessionID, quizID, &result)
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	r := response.NewSuccess("تم إرسال النتيجة بنجاح")
	c.JSON(http.StatusCreated, r)
}

func getIdsAsUint(c *gin.Context) (uint, uint, error) {
	sessionID := c.Param("session_id")
	quizID := c.Param("quiz_id")
	sessionIDInt, err := strconv.Atoi(sessionID)
	if err != nil {
		return 0, 0, err
	}
	quizIDInt, err := strconv.Atoi(quizID)
	if err != nil {
		return 0, 0, err
	}
	return uint(sessionIDInt), uint(quizIDInt), nil
}
