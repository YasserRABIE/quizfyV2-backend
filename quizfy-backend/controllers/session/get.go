package session

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/YasserRABIE/QUIZFYv2/migrations/session_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/models/session"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	quizID := c.Param("quiz_id")
	quizIDInt, err := strconv.Atoi(quizID)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	s, err := session_migrations.Get(userID, uint(quizIDInt))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r := response.NewSuccess(&session.Session{Status: "not_started"})
			c.JSON(http.StatusOK, r)
			return
		}
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	r := response.NewSuccess(s)
	c.JSON(http.StatusOK, r)
}
