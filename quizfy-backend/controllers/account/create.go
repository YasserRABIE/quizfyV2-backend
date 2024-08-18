package user

import (
	"errors"
	"net/http"

	user_migrations "github.com/YasserRABIE/QUIZFYv2/migrations/account_migrations"
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/YasserRABIE/QUIZFYv2/models/user"
	"github.com/YasserRABIE/QUIZFYv2/services/auth"
	"github.com/YasserRABIE/QUIZFYv2/utils"
	"github.com/gin-gonic/gin"
)

const (
	errCreateAccount = "لم يتم إنشاء الحساب بنجاح!، يرجى المحاولة مرة أخرى"
)

func Create(c *gin.Context) {
	var a user.Account

	// Bind request body
	err := c.ShouldBindJSON(&a)
	if err != nil {
		err = errors.New("يرجى ملأ كل الحقول!")
		utils.HandleError(c, err, http.StatusBadRequest)
		return
	}

	// hash password
	a.Password, err = auth.HashPass(a.Password)
	if err != nil {
		utils.HandleError(c, errors.New(errCreateAccount), http.StatusInternalServerError)
		return
	}

	// Create student in database
	err = user_migrations.Create(&a)
	if err != nil {
		utils.HandleError(c, errors.New(errCreateAccount), http.StatusBadRequest)
		return
	}

	// Generate token
	token, err := auth.CreateToken(a.ID)
	if err != nil {
		utils.HandleError(c, errors.New(errCreateAccount), http.StatusInternalServerError)
		return
	}

	// Send success response
	r := response.NewSuccess(token)
	c.JSON(http.StatusCreated, r)
}
