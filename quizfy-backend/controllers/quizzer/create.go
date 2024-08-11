package quizzer

import (
	"net/http"

	quizzerM "github.com/YasserRABIE/QUIZFYv2/migrations/quizzer_migration"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var q user.Quizzer

	// Bind request body
	err := c.ShouldBindJSON(&q)
	if utils.HandleError(c, err, http.StatusBadRequest) {
		return
	}

	// Create quizzer in database
	err = quizzerM.Create(&q)
	if utils.HandleError(c, err, http.StatusBadRequest) {
		return
	}

	// Generate token
	token, err := auth.CreateToken(q.ID)
	if utils.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	// Send success response
	r := response.NewSuccess(token)
	c.JSON(http.StatusCreated, r)
}
