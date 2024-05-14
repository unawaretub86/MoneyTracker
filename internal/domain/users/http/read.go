package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unawaretub86/MoneyTracker/utils"
)

func (userHandler Handler) GetUsers(c *gin.Context) {
	users, err := userHandler.UseUser.GetUsers()
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixUser, users)
}

func (userHandler Handler) GetUserByID(c *gin.Context) {
	var idParam uint

	if err := c.ShouldBindUri(&idParam); err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	user, err := userHandler.UseUser.GetUserByID(idParam)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	utils.EndWithStatus(c, http.StatusOK, suffixUser, user)
}