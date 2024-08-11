package quizzer

import (
	"errors"
	"net/http"

	"github.com/YasserRABIE/QUIZFYv2/db"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context) {
	var quizzer user.Quizzer

	// Bind request body
	err := c.ShouldBindJSON(&quizzer)
	if utils.HandleError(c, err, http.StatusBadRequest) {
		return
	}

	// Create quizzer
	err = createQuizzer(&quizzer)
	if utils.HandleError(c, err, http.StatusBadRequest) {
		return
	}

	// Generate token
	token, err := auth.CreateToken(quizzer.Phone)
	if utils.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	// Send success response
	r := response.NewSuccess(token)
	c.JSON(http.StatusCreated, r)
}

func createQuizzer(q *user.Quizzer) error {
	err := db.Conn.Create(q).Error
	if err != nil {
		// Check if the error is due to a unique constraint violation
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("الرقم ده مسجل قبل كده")
		}
		return errors.New("حصل خطأ في السيرفر. حاول تاني بعد شوية")
	}
	return nil
}
