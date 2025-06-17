package main

import (
	"log"
	"temporal-poc/internal/common/constants"
	"temporal-poc/internal/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	dependencies, err := InitDependencies()
	if err != nil {
		log.Fatalln("Unable to initialize dependencies.", err)
		return
	}

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
		return
	}
	defer c.Close()

	w := worker.New(c, constants.UpdateProfilePictureWorkflowId, worker.Options{})
	w.RegisterWorkflow(workflow.UpdateProfilePictureWorkflow)
	w.RegisterActivity(dependencies.UserActivity.GetUser)
	w.RegisterActivity(dependencies.UserActivity.UpdateUser)
	w.RegisterActivity(dependencies.UserActivity.SendNotification)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
