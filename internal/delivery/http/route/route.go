package route

import (
	"github.com/gin-gonic/gin"
	"temporal-poc/internal/delivery/http"
)

type RouteConfig struct {
	R                     *gin.Engine
	HealthcheckController *http.HealthcheckController
	UserController        *http.UserController
}

func (c *RouteConfig) SetupGuestRoute() {
	c.R.GET("/ping", c.HealthcheckController.Ping)
	c.R.POST("/user/profile-picture", c.UserController.UpdateProfilePicture)
	c.R.POST("/user/profile-picture/workflow", c.UserController.UpdateProfilePictureUsingWorkflow)
}
