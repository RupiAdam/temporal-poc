package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"temporal-poc/internal/usecase"
)

type UserController struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserController(usecase *usecase.UserUsecase) *UserController {
	return &UserController{
		UserUsecase: usecase,
	}
}

func (c *UserController) UpdateProfilePicture(ctx *gin.Context) {
	// Logic to update the images picture using UserUsecase
	// This is a placeholder for the actual implementation

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.UserUsecase.UpdateProfilePicture(ctx, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar"})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Profile picture has been updated successfully",
	})
	ctx.Status(200)

}
