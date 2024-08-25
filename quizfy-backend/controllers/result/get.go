package result

import (
	"net/http"
	"strconv"

	"github.com/YasserRABIE/QUIZFYv2/migrations/result_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	sessionID := c.Param("session_id")
	sessionIDInt, err := strconv.Atoi(sessionID)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	result, err := result_migrations.Get(uint(sessionIDInt))
	if err != nil {
		utils.HandleError(c, err, http.StatusNotFound)
		return
	}

	r := response.NewSuccess(result)
	c.JSON(http.StatusOK, r)
}
