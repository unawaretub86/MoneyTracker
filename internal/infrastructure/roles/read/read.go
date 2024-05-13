package read

import (
	"github.com/gin-gonic/gin"

	yourHandler "github.com/unawaretub86/MoneyTracker/internal/domain/money-tracker/http"
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
)

type read struct {
	container *dependencies.Container
}

func NewRead(container *dependencies.Container) *read {
	return &read{
		container: container,
	}
}

func (read *read) RegisterRoutes(basePath string, r *gin.Engine) {
	handler := yourHandler.NewHandler(read.container)

	r.GET("/ping", handler.Ping)
}
