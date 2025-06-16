package main

import (
	"log"
	"temporal-poc/internal/activity"
	"temporal-poc/internal/common/constants"
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

	w := worker.New(c, constants.UpdateProfilePictureWorkflowId, worker.Options{})
	userRepository := repository.NewUserRepository()
	userActivity := activity.NewUserActivity(userRepository)

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(workflow.UpdateProfilePictureWorkflow)
	w.RegisterActivity(userActivity.GetUser)
	w.RegisterActivity(userActivity.UpdateUser)
	w.RegisterActivity(userActivity.SendNotification)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
