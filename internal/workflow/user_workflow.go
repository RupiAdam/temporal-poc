package workflow

import (
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"log"
	"temporal-poc/internal/activity"
	"temporal-poc/internal/model"
	"temporal-poc/internal/repository"
	"time"
)

func UpdateProfilePictureWorkflow(ctx workflow.Context) (string, error) {

	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        10, // 0 is unlimited retries
		NonRetryableErrorTypes: []string{"ImageProcessError", "InvalidFileError"},
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: 2 * time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retryPolicy,
	}

	// Apply the options.
	ctx = workflow.WithActivityOptions(ctx, options)

	userRepository := repository.NewUserRepository()
	userActivity := activity.NewUserActivity(userRepository)
	var userData *model.UserModel

	getUserErr := workflow.ExecuteActivity(ctx, userActivity.GetUser).Get(ctx, &userData)
	if getUserErr != nil {
		// this is an error that cannot be retried, so we log it and return
		log.Printf("Error getting user data: %v", getUserErr)
		return "", getUserErr
	}

	errUpdateProfilePicture := workflow.ExecuteActivity(ctx, userActivity.UpdateUser, userData).Get(ctx, nil)
	if errUpdateProfilePicture != nil {
		// this is an error that cannot be retried, so we log it and return
		log.Printf("Error updating profile picture: %v", errUpdateProfilePicture)
		return "", errUpdateProfilePicture
	}

	sendNotificationErr := workflow.ExecuteActivity(ctx, userActivity.SendNotification, userData).Get(ctx, nil)
	if sendNotificationErr != nil {
		log.Printf("Error sending notification: %v", sendNotificationErr)
		return "", sendNotificationErr
	}

	return "Workflow finished", nil
}
