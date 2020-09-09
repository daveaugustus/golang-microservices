package controllers

import (
	"net/http"
	"strconv"

	"github.com/davetweetlive/golang-microservices/mvc/services"
	"github.com/davetweetlive/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userIdPram := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdPram, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "User id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
