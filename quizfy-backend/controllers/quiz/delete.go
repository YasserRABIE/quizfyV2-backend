package quiz

import (
	"net/http"
	"strconv"

	"github.com/YasserRABIE/QUIZFYv2/migrations/quiz_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	err = quiz_migrations.DeleteByID(uint(idInt))
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	r := response.NewSuccess("تم حذف الاختبار بنجاح")
	c.JSON(http.StatusOK, r)
}
