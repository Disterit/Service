package repository

import (
	"Service/grpc/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	createUserQuery = `INSERT INTO users (username, password) VALUES ($1, $2)`
)

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) User {
	return &userRepository{pool: pool}
}

func (u *userRepository) Register(ctx context.Context, user models.Users) error {
	_, err := u.pool.Exec(ctx, createUserQuery, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}
