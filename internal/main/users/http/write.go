package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/MoneyTracker/internal/main/users/http/dto"
	"github.com/unawaretub86/MoneyTracker/utils"
)

func (userHandler Handler) UpdateUser(c *gin.Context) {
	req := dto.UpdateUserRequest{}

	id, err := utils.GetHttpID(c)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	if err := c.ShouldBindJSON(req); err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	user, err := userHandler.UseUser.UpdateUser(uint(*id), &req)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	utils.EndWithStatus(c, http.StatusCreated, user)
}
