package repository

import (
	"Service/grpc/internal/models"
	"context"
)

type User interface {
	Register(ctx context.Context, user models.Users) error
}

type Repository struct {
	User User
}

func NewRepository(user User) *Repository {
	return &Repository{
		User: user,
	}
}
