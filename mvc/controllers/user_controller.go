package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muchlist/golang-microservices/mvc/services"
	"github.com/muchlist/golang-microservices/mvc/utils"
)

//GetUser mengembalikan user
func GetUser(c *gin.Context) {
	//c.Param for parameter  localhost/users/user_id
	//c.Query for query localhost/users?user_id=
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "UserID harus berupa angka!",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		//c.JSON(apiErr.StatusCode, apiErr) <- normalnya seperti ini
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userID)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
