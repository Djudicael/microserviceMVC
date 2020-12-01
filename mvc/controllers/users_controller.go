package controllers

import (
	"net/http"
	"strconv"

	"github.com/Djudicael/microserviceMVC/mvc/services"
	"github.com/Djudicael/microserviceMVC/mvc/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	userId, err := (strconv.ParseInt(c.Param("user_id"), 10, 64))

	if err != nil {
		// Just return the bad resquest to the client
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		//c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		// c.JSON(apiErr.StatusCode, apiErr)
		utils.RespondError(c, apiErr)
		return
	}

	//return user to client
	// c.JSON(http.StatusOK, user)
	utils.Respond(c, http.StatusOK, user)
}
