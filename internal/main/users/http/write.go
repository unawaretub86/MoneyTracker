package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/MoneyTracker/internal/main/users/http/dto"
	"github.com/unawaretub86/MoneyTracker/utils"
)

func (userHandler Handler) CreateUser(c *gin.Context) {
	req := dto.CreateUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	wasCreated, err := userHandler.UseUser.CreateUser(req)
	if err != nil {
		utils.EndWithStatusError(c, http.StatusBadRequest, suffixErr, err)
		return
	}

	utils.EndWithStatus(c, http.StatusCreated, wasCreated)
}
