package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/unawaretub86/MoneyTracker/utils"
)

func (userHandler Handler) GetUsers(c *gin.Context) {
	users, err := userHandler.UseUser.GetUsers()
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, users)
}

func (userHandler Handler) GetUserByID(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	user, err := userHandler.UseUser.GetUserByID(uint(id))
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, user)
}
