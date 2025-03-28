package repository

import (
	"Service/grpc/internal/models"
	"context"
)

type User interface {
	Register(ctx context.Context, user models.Users) error
	CheckExists(ctx context.Context, user models.Users) bool
}

type Repository struct {
	User User
}

func NewRepository(user User) *Repository {
	return &Repository{
		User: user,
	}
}
