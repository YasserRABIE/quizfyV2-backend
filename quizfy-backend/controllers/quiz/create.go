package quiz

import (
	"errors"
	"net/http"

	"github.com/YasserRABIE/QUIZFYv2/migrations/quiz_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var quizData quiz.Quiz

	// add user_id to quizData from token
	u, _ := c.Get("user")

	quizData.UserID = u.(*user.Account).ID

	// Bind request body
	err := c.ShouldBindJSON(&quizData)
	if err != nil {
		err = errors.New("يرجى ملأ كل الحقول")
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}
	// Set teacher name
	quizData.TeacherName = u.(*user.Account).Name

	// Create quiz in database
	err = quiz_migrations.Create(&quizData)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// Send success response
	r := response.NewSuccess(quizData)
	c.JSON(http.StatusCreated, r)
}
