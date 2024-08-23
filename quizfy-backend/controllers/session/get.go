package session

import (
	"net/http"
	"strconv"

	"github.com/YasserRABIE/QUIZFYv2/migrations/session_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	quizID := c.Param("quiz_id")
	quizIDInt, err := strconv.Atoi(quizID)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	session, err := session_migrations.Get(userID, uint(quizIDInt))
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	r := response.NewSuccess(session)
	c.JSON(http.StatusOK, r)
}
