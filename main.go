package main

import (
	"github.com/gin-gonic/gin"
	"temporal-poc/internal/delivery/http"
	"temporal-poc/internal/delivery/http/route"
	"temporal-poc/internal/usecase"
)

func main() {
	imageProcessingUsecase := usecase.NewImageProcessingUsecase()
	userUsecase := usecase.NewUserUsecase(imageProcessingUsecase)

	healthcheckController := http.NewHealthcheckController()
	userController := http.NewUserController(userUsecase)
	r := gin.Default()

	routeConfig := &route.RouteConfig{
		R:                     r,
		HealthcheckController: healthcheckController,
		UserController:        userController,
	}
	routeConfig.SetupGuestRoute()

	r.Run(":8080")
}
