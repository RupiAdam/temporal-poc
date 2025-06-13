package http

import "github.com/gin-gonic/gin"

type HealthcheckController struct{}

func NewHealthcheckController() *HealthcheckController {
	return &HealthcheckController{}
}

func (c *HealthcheckController) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":  "ok",
		"message": "pong",
	})
	ctx.Status(200)
}
