package route

import (
	"github.com/gin-gonic/gin"
	"temporal-poc/internal/delivery/http"
)

type RouteConfig struct {
	R                     *gin.Engine
	HealthcheckController *http.HealthcheckController
}

func (c *RouteConfig) SetupGuestRoute() {
	c.R.GET("/ping", c.HealthcheckController.Ping)
}
