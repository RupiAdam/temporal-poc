package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"temporal-poc/internal/config"
)

func main() {
	viper := config.NewViper()

	r := gin.Default()
	r.Use(requestid.New())
	r.Use(gin.Recovery())

	dep, err := InitDependencies(r)
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
		return
	}

	dep.Route.SetupGuestRoute()

	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))
	err = r.Run(port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
