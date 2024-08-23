package session

import (
	"net/http"
	"strconv"

	"github.com/YasserRABIE/QUIZFYv2/migrations/session_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	userIdAny, _ := c.Get("user_id")
	userID := userIdAny.(uint)
	quizID := c.Param("quiz_id")
	quizIDInt, err := strconv.Atoi(quizID)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// Create a new session
	session, err := session_migrations.Create(uint(quizIDInt), userID)
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	// Send success response
	r := response.NewSuccess(session)
	c.JSON(http.StatusCreated, r)
}
