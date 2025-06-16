package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.temporal.io/sdk/client"
	"log"
	"net/http"
	"temporal-poc/internal/common/constants"
	"temporal-poc/internal/helper"
	"temporal-poc/internal/usecase"
	"temporal-poc/internal/workflow"
)

type UserController struct {
	ResponseHelper helper.ResponseHelper
	UserUsecase    *usecase.UserUsecase
}

func NewUserController(usecase *usecase.UserUsecase) *UserController {
	return &UserController{
		ResponseHelper: helper.ResponseHelper{},
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

func (c *UserController) UpdateProfilePictureUsingWorkflow(ctx *gin.Context) {
	// Logic to update the images picture using UserUsecase with Temporal Workflow
	// This is a placeholder for the actual implementation

	_, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, c.ResponseHelper.GenerateError(ctx, "File is required"))
		return
	}

	// Create the client object just once per process
	temporalClient, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer temporalClient.Close()

	options := client.StartWorkflowOptions{
		TaskQueue: constants.UpdateProfilePictureWorkflowId,
	}

	log.Printf("Starting update profile picture")

	we, err := temporalClient.ExecuteWorkflow(context.Background(), options, workflow.UpdateProfilePictureWorkflow)
	if err != nil {
		log.Fatalln("Unable to start the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	ctx.JSON(200, gin.H{
		"status":      "success",
		"message":     "Profile picture process has been added to workflow",
		"workflow_id": we.GetID(),
		"run_id":      we.GetRunID(),
	})
	ctx.Status(200)
}
