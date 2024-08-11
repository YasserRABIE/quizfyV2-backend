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

func Get(c *gin.Context) {
	var q user.QLoginReq

	// Bind request body
	err := c.ShouldBindJSON(&q)
	if utils.HandleError(c, err, http.StatusBadRequest) {
		return
	}

	// Get quizzer from database
	quizzer, err := quizzerM.Get(q.Phone)
	if utils.HandleError(c, err, http.StatusBadRequest) {
		return
	}

	// validate password with the hashed password from the database
	err = auth.ValidatePassword(quizzer.Password, q.Password)
	if utils.HandleError(c, err, http.StatusUnauthorized) {
		return
	}

	// Generate token
	token, err := auth.CreateToken(quizzer.ID)
	if utils.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	// Send success response
	r := response.NewSuccess(token)
	c.JSON(http.StatusOK, r)
}
