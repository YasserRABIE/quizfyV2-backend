package session

import (
	"net/http"
	"strconv"

	"github.com/YasserRABIE/QUIZFYv2/migrations/session_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	sessionID := c.Param("session_id")
	sessionIDInt, err := strconv.Atoi(sessionID)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	session, err := session_migrations.Update(uint(sessionIDInt))
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	r := response.NewSuccess(session)
	c.JSON(http.StatusOK, r)
}
