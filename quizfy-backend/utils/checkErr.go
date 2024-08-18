package utils

import (
	"github.com/YasserRABIE/QUIZFYv2/models/response"
	"github.com/gin-gonic/gin"
)

/*
HandleError is a utility function to handle errors

and return a response with the error message
*/
func HandleError(c *gin.Context, err error, statusCode int) {
	if err != nil {
		r := response.NewError(err.Error())
		c.JSON(statusCode, r)
		c.Abort()
	}
}
