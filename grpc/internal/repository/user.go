package repository

import (
	"Service/grpc/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	pool *pgxpool.Pool
}

func newUserRepository(pool *pgxpool.Pool) User {
	return &userRepository{pool: pool}
}

func (u *userRepository) Register(ctx context.Context, user models.Users) error {

	return nil
}
