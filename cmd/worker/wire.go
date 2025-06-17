//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"temporal-poc/internal/activity"
	"temporal-poc/internal/repository"
)

type Dependencies struct {
	UserActivity *activity.UserActivity
}

func InitDependencies() (*Dependencies, error) {
	wire.Build(
		repository.NewUserRepository,
		activity.NewUserActivity,
		wire.Struct(new(Dependencies), "*"),
	)

	return &Dependencies{}, nil
}
