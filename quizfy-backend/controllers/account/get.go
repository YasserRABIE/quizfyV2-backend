package user

import (
	"errors"
	"fmt"
	"net/http"

	user_migrations "github.com/YasserRABIE/QUIZFYv2/migrations/account_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var req user.LoginReq

	// Bind request body
	err := c.ShouldBindJSON(&req)
	if err != nil {
		err = errors.New("يرجى ملأ كل الحقول!")
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}
	fmt.Println(req)
	// Get quizzer from database
	quizzer, err := user_migrations.GetByPhone(req.Phone)
	if err != nil {
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// validate password with the hashed password from the database
	err = auth.ValidatePassword(quizzer.Password, req.Password)
	if err != nil {
		err = errors.New("كلمة المرور غير صحيحة!")
		utils.HandleError(c, err, http.StatusUnauthorized)
		return
	}

	// Generate token
	token, err := auth.CreateToken(quizzer.ID)
	if err != nil {
		err = errors.New("حدث خطأ ما، يرجى المحاولة مرة أخرى!")
		utils.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	// Send success response
	r := response.NewSuccess(token)
	c.JSON(http.StatusOK, r)
}

func GetValidatedUser(c *gin.Context) {
	// Get user data from context
	user_data, exists := c.Get("user")
	if !exists {
		err := errors.New("غير مصرح")
		utils.HandleError(c, err, http.StatusUnauthorized)
		return
	}

	// Send success response
	r := response.NewSuccess(user_data.(*user.Account))
	c.JSON(http.StatusOK, r)
}
