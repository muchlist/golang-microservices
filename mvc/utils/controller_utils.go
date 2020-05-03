package utils

import "github.com/gin-gonic/gin"

//Respond akan membaca header, secara default akan mengembalikan JSON
func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

//RespondError mengembalikan error code sesuai struct Application error
func RespondError(c *gin.Context, err *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, body)
		return
	}
	c.JSON(err.StatusCode, body)
}
