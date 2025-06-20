// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"temporal-poc/internal/activity"
	"temporal-poc/internal/repository"
)

// Injectors from wire.go:

func InitDependencies() (*Dependencies, error) {
	userRepository := repository.NewUserRepository()
	userActivity := activity.NewUserActivity(userRepository)
	dependencies := &Dependencies{
		UserActivity: userActivity,
	}
	return dependencies, nil
}

// wire.go:

type Dependencies struct {
	UserActivity *activity.UserActivity
}
