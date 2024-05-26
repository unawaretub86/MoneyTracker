package read

import (
	"github.com/gin-gonic/gin"

	moneyHandler "github.com/unawaretub86/MoneyTracker/internal/main/money-tracker/http"
	userHandler "github.com/unawaretub86/MoneyTracker/internal/main/users/http"
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

	userHandler := userHandler.NewUserHandler(read.container)
	moneyHandler := moneyHandler.NewHandler(read.container)

	v1Group := r.Group(basePath + "/v1")

	v1Group.GET("/users", userHandler.GetUsers)
	v1Group.GET("/user/:id", userHandler.GetUserByID)

	r.GET("/ping", moneyHandler.Ping)
}
