package main

import (
	"fmt"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"log"
	"temporal-poc/internal/config"
	"temporal-poc/internal/delivery/http"
	"temporal-poc/internal/delivery/http/route"
	"temporal-poc/internal/helper"
	"temporal-poc/internal/usecase"
)

func main() {
	viper := config.NewViper()
	imageProcessingUsecase := helper.NewImageProcessingHelper()
	userUsecase := usecase.NewUserUsecase(imageProcessingUsecase)

	healthcheckController := http.NewHealthcheckController()
	userController := http.NewUserController(userUsecase)
	r := gin.Default()
	r.Use(requestid.New())

	routeConfig := &route.RouteConfig{
		R:                     r,
		HealthcheckController: healthcheckController,
		UserController:        userController,
	}
	routeConfig.SetupGuestRoute()

	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))
	err := r.Run(port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
