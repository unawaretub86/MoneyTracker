package write

import (
	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
)

type write struct {
	container *dependencies.Container
}

func NewWrite(container *dependencies.Container) *write {
	return &write{
		container: container,
	}
}

func (write *write) RegisterRoutes(basePath string, r *gin.Engine) {
}
