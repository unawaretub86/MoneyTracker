package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func errorResponse(err error, suffix string) gin.H {
	return gin.H{suffix: err.Error()}
}

func EndWithStatus(c *gin.Context, status int, body any) {
	c.JSON(status, body)
}

func EndWithStatusError(c *gin.Context, status int, suffix string, err error) {
	c.JSON(status, errorResponse(err, suffix))
}

func GetHttpID(c *gin.Context) (*uint64,error) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, err
	}

	return &id, nil
}