package activity

import (
	"context"
	"go.temporal.io/sdk/activity"
	"temporal-poc/internal/model"
	"temporal-poc/internal/repository"
)

type UserActivity struct {
	UserRepository *repository.UserRepository
}

func NewUserActivity(userRepository *repository.UserRepository) *UserActivity {
	return &UserActivity{
		UserRepository: userRepository,
	}
}

func (c *UserActivity) GetUser(ctx context.Context) (*model.UserModel, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Starting GetUser activity")

	data, err := c.UserRepository.Get()
	if err != nil {
		logger.Error("Error getting user data", "error", err)
		return nil, err
	}

	logger.Info("GetUser activity completed successfully", "user", data)
	return data, nil
}

func (c *UserActivity) UpdateUser(ctx context.Context, user *model.UserModel) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Starting UpdateUser activity", "user", user)

	err := c.UserRepository.Update()
	if err != nil {
		logger.Error("Error updating user data", "error", err)
		return err
	}

	logger.Info("UpdateUser activity completed successfully", "user", user)
	return nil
}

func (c *UserActivity) SendNotification(ctx context.Context, user *model.UserModel) (map[string]interface{}, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Starting SendNotification activity", "user", user)

	resp, err := repository.SendNotification()
	if err != nil {
		logger.Error("Error sending notification", "error", err)
		return nil, err
	}

	logger.Info("SendNotification activity completed successfully", "response", resp)
	return resp, nil
}
