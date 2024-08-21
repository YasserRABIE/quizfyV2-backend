package question

import (
	"errors"
	"net/http"

	question_migrations "github.com/YasserRABIE/QUIZFYv2/migrations/questions_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r quiz.QuestionReq
	if err := c.ShouldBindJSON(&r); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// Set the QuestionID for each option
	for i := range r.Options {
		r.Options[i].QuestionID = r.ID
	}

	if err := question_migrations.Create(&r.Question); err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	
	if r.ImageData != nil {
		// Upload the image to the server
		r.ImagePath = utils.UploadImage(
			r.ImageData.Image,
			r.ImageData.Extension,
			r.QuizID,
			r.ID,
		)
		if r.ImagePath == "" {
			utils.HandleError(c, errors.New("failed to upload image"), http.StatusInternalServerError)
			return
		}
		// Update the question with the image paths
		if err := question_migrations.UpdateImage(
			r.ID,
			r.ImagePath,
		); err != nil {
			utils.HandleError(c, err, http.StatusInternalServerError)
			return
		}
	}
	res := response.NewSuccess(r.Question)
	c.JSON(http.StatusCreated, res)
}
