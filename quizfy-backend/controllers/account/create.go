package user

import (
	"net/http"

	user_migrations "github.com/YasserRABIE/QUIZFYv2/migrations/account_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var a user.Account

	// Bind request body
	err := c.ShouldBindJSON(&a)
	if utils.HandleError(c, err, http.StatusBadRequest) {
		return
	}

	// hash password
	a.Password, err = auth.HashPass(a.Password)
	if utils.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	// Create student in database
	err = user_migrations.Create(&a)
	if utils.HandleError(c, err, http.StatusBadRequest) {
		return
	}

	// Generate token
	token, err := auth.CreateToken(a.ID)
	if utils.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	// Send success response
	r := response.NewSuccess(token)
	c.JSON(http.StatusCreated, r)
}
