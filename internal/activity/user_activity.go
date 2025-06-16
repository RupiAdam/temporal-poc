package activity

import (
	"context"
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
	data, err := c.UserRepository.Get()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *UserActivity) UpdateUser(ctx context.Context, user *model.UserModel) error {
	err := c.UserRepository.Update()
	if err != nil {
		return err
	}

	return nil
}

func (c *UserActivity) SendNotification(ctx context.Context, user *model.UserModel) (map[string]interface{}, error) {
	resp, err := repository.SendNotification()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
