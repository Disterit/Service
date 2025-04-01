package repository

import (
	"Service/grpc/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	createUserQuery  = `INSERT INTO users (username, password_hash, email) VALUES ($1, $2, $3)`
	loginQuery       = `SELECT password_hash, is_active FROM users WHERE username = $1`
	checkExistsQuery = `SELECT username FROM users WHERE username = $1 OR email = $2`
)

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) User {
	return &userRepository{pool: pool}
}

func (u *userRepository) Register(ctx context.Context, user models.Users) error {
	_, err := u.pool.Exec(ctx, createUserQuery, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) CheckExists(ctx context.Context, user models.Users) bool {
	var username string
	_ = u.pool.QueryRow(ctx, checkExistsQuery, user.Username, user.Email).Scan(&username)
	if username == "" {
		return true
	}

	return false
}

func (u *userRepository) Login(ctx context.Context, user models.Users) (models.Users, error) {
	var userOut models.Users
	err := u.pool.QueryRow(ctx, loginQuery, user.Username).Scan(&userOut.Password, &userOut.Active)
	if err != nil {
		return userOut, err
	}

	return userOut, nil
}
