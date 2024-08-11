package student

import (
	"net/http"

	studentM "github.com/YasserRABIE/QUIZFYv2/migrations/student_migration"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var s user.Student

	// Bind request body
	err := c.ShouldBindJSON(&s)
	if utils.HandleError(c, err, http.StatusBadRequest) {
		return
	}

	// hash password
	s.Password, err = auth.HashPass(s.Password)
	if utils.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	// Create student in database
	err = studentM.Create(&s)
	if utils.HandleError(c, err, http.StatusBadRequest) {
		return
	}

	// Generate token
	token, err := auth.CreateToken(s.ID)
	if utils.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	// Send success response
	r := response.NewSuccess(token)
	c.JSON(http.StatusCreated, r)
}
