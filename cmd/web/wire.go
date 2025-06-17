//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"temporal-poc/internal/delivery/http"
	"temporal-poc/internal/delivery/http/route"
	"temporal-poc/internal/helper"
	"temporal-poc/internal/usecase"
)

type Dependencies struct {
	Route                 *route.RouteConfig
	HealthcheckController *http.HealthcheckController
	UserController        *http.UserController
}

func InitDependencies(r *gin.Engine) (*Dependencies, error) {
	wire.Build(
		helper.NewImageProcessingHelper,

		usecase.NewUserUsecase,

		http.NewHealthcheckController,
		http.NewUserController,

		route.NewRouteConfig,

		wire.Struct(new(Dependencies), "*"),
	)
	return &Dependencies{}, nil
}
