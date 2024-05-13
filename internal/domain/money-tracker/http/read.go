package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unawaretub86/MoneyTracker/utils"
)

func (handler Handler) Ping(c *gin.Context) {
	utils.EndWithStatus(c, http.StatusOK, "response", "pong")
}
