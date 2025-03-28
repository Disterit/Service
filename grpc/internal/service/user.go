package service

import (
	"Service/grpc/internal/models"
	"Service/grpc/internal/repository"
	"context"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.User
	log  *zap.SugaredLogger
}

func NewUserService(repo repository.User, log *zap.SugaredLogger) User {
	return &userService{
		repo: repo,
		log:  log,
	}
}

func (u *userService) Register(ctx context.Context, user models.Users) error {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Errorw("error to hash password", "error", err)
		return err
	}
	user.Password = string(newPassword)

	err = u.repo.Register(ctx, user)
	if err != nil {
		u.log.Errorw("error to create user in db", "error", err)
		return err
	}

	return nil
}
