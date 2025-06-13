package utilities

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type ResponseHelper struct{}

func (c *ResponseHelper) GenerateError(ctx *gin.Context, message string) gin.H {
	return gin.H{
		"request-id": requestid.Get(ctx),
		"message":    message,
	}
}
