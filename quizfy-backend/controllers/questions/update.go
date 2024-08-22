package question

import (
	"net/http"
	"os"

	question_migrations "github.com/YasserRABIE/QUIZFYv2/migrations/questions_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/quiz"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	var r quiz.QuestionReq
	if err := c.ShouldBindJSON(&r); err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	if err := handleImageUpdate(&r, r.Question.ID); err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	// Update the question options and question itself
	if err := updateQuestionAndOptions(&r, r.Question.ID); err != nil {
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	res := response.NewSuccess("تم تعديل السؤال بنجاح")
	c.JSON(http.StatusOK, res)
}

func handleImageUpdate(r *quiz.QuestionReq, questionID uint) error {
	// if there is an image and the user wants to remove it
	if r.ImagePath != "" && r.NoImage {
		if _, err := os.Stat(r.ImagePath); err == nil {
			if err := os.Remove(r.ImagePath); err != nil {
				return err
			}
		}
		r.ImagePath = ""
	}

	// if there is an image and the user wants to update it
	if r.ImageData != nil {
		imagePath, err := utils.UploadImage(r.ImageData.Image, r.ImageData.Extension, r.QuizID, questionID)
		if err != nil {
			return err
		}
		r.ImagePath = imagePath
	}
	return nil
}

func updateQuestionAndOptions(r *quiz.QuestionReq, questionID uint) error {
	var ids []uint
	for i := range r.Options {
		r.Options[i].QuestionID = questionID
		if r.Options[i].ID != 0 {
			ids = append(ids, r.Options[i].ID)
		}
	}

	err := question_migrations.DeleteOptions(questionID, ids)
	if err != nil {
		return err
	}

	return question_migrations.Update(&r.Question)
}
