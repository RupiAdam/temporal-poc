package main

import (
	"github.com/gin-gonic/gin"
	"temporal-poc/internal/delivery/http"
	"temporal-poc/internal/delivery/http/route"
)

func main() {
	healthcheckController := http.NewHealthcheckController()
	r := gin.Default()

	routeConfig := &route.RouteConfig{
		R:                     r,
		HealthcheckController: healthcheckController,
	}
	routeConfig.SetupGuestRoute()

	r.Run(":8080")
}
