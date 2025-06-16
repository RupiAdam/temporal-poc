package repository

import (
	"fmt"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"math/rand"
	"temporal-poc/internal/model"
	"time"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (c *UserRepository) Get(ctx *gin.Context) (*model.UserModel, error) {
	// Random seed for better randomness each run
	rand.Seed(time.Now().UnixNano())

	// Random delay between 1s to 2m
	delay := time.Duration(1+rand.Intn(120)) * time.Second
	time.Sleep(delay)

	// Simulate 40% failure rate
	if rand.Float64() < 0.4 {
		return nil, fmt.Errorf("[%v] intermittent error occurred after %v delay", requestid.Get(ctx), delay)
	}

	fmt.Printf("[%v] Task completed successfully after %v delay\n", requestid.Get(ctx), delay)
	user := &model.UserModel{
		Id:             "12345",
		Name:           "John Doe",
		Email:          "mail@mail.com",
		ProfilePicture: "https://example.com/profile.jpg",
	}
	return user, nil
}

func (c *UserRepository) Update(ctx *gin.Context) error {
	// Random seed for better randomness each run
	rand.Seed(time.Now().UnixNano())

	// Random delay between 1s to 2m
	delay := time.Duration(1+rand.Intn(120)) * time.Second
	time.Sleep(delay)

	// Simulate 40% failure rate
	if rand.Float64() < 0.4 {
		return fmt.Errorf("[%v] intermittent error occurred after %v delay", requestid.Get(ctx), delay)
	}

	fmt.Printf("[%v] Task completed successfully after %v delay\n", requestid.Get(ctx), delay)
	return nil
}
