package main

import (
	"log"
	"temporal-poc/internal/common/constants"
	"temporal-poc/internal/helper"
	"temporal-poc/internal/repository"
	"temporal-poc/internal/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, constants.UPDATE_PROFILE_PICTURE_WORKFLOW_ID, worker.Options{})
	imageProcessingHelper := helper.NewImageProcessingHelper()
	userRepository := repository.NewUserRepository()

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(workflow.UpdateProfilePictureWorkflow)
	w.RegisterActivity(imageProcessingHelper.Resize)
	w.RegisterActivity(userRepository.Get)
	w.RegisterActivity(userRepository.Update)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
