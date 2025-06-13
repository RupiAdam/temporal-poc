package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"temporal-poc/internal/usecase"
	"temporal-poc/internal/utilities"
)

type UserController struct {
	ResponseHelper utilities.ResponseHelper
	UserUsecase    *usecase.UserUsecase
}

func NewUserController(usecase *usecase.UserUsecase) *UserController {
	return &UserController{
		ResponseHelper: utilities.ResponseHelper{},
		UserUsecase:    usecase,
	}
}

func (c *UserController) UpdateProfilePicture(ctx *gin.Context) {
	// Logic to update the images picture using UserUsecase
	// This is a placeholder for the actual implementation

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, c.ResponseHelper.GenerateError(ctx, "File is required"))
		return
	}

	err = c.UserUsecase.UpdateProfilePicture(ctx, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, c.ResponseHelper.GenerateError(ctx, err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": "Profile picture has been updated successfully",
	})
	ctx.Status(200)

}
