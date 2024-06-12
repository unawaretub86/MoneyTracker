package write

import (
	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
	userHandler "github.com/unawaretub86/MoneyTracker/internal/main/users/http"
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
	
	userHandler := userHandler.NewUserHandler(write.container)

	v1Group := r.Group(basePath + "/v1")

	v1Group.PATCH("/users/:id", userHandler.UpdateUser)
}
