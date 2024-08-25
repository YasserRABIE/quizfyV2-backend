package answer

import (
	"net/http"

	"github.com/YasserRABIE/QUIZFYv2/migrations/answer_migrations"
	"github.com/YasserRABIE/QUIZFYv2/migrations/session_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/models/result"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var answers []result.Answer
	if err := c.ShouldBindJSON(&answers); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	_, err := session_migrations.Finish(answers[0].SessionID)
	if err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	if err := answer_migrations.Create(answers); err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	r := response.NewSuccess(answers)
	c.JSON(http.StatusCreated, r)
}
