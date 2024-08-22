package question

import (
	"net/http"
	"os"
	"strconv"

	question_migrations "github.com/YasserRABIE/QUIZFYv2/migrations/questions_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	questionID := c.Param("question_id")
	questionIDInt, err := strconv.Atoi(questionID)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	imagePath, err := question_migrations.DeleteByID(uint(questionIDInt))
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}
	// delete question image if exists
	if imagePath != "" {
		err = os.Remove(imagePath)
		if err != nil {
			utils.HandleError(c, err, http.StatusBadRequest)
			return
		}
	}

	r := response.NewSuccess("تم حذف السؤال بنجاح")
	c.JSON(http.StatusOK, r)
}
