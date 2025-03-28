package service

import (
	"Service/grpc/internal/models"
	"Service/grpc/internal/repository"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
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

	ok := u.repo.CheckExists(ctx, user)
	if !ok {
		u.log.Info("User already exists")
		return errors.New("user already created")
	}

	user.Password = generateHashPassword(user.Password)

	err := u.repo.Register(ctx, user)
	if err != nil {
		u.log.Errorw("error to create user in db", "error", err)
		return err
	}

	return nil
}

func (u *userService) Login(ctx context.Context, user models.Users) (string, error) {

	user.Password = generateHashPassword(user.Password)

	active, err := u.repo.Login(ctx, user)
	if err != nil {
		u.log.Errorw("error to login", "error", err)
		return "", err
	}

	// тут логика создание токена будет если пользователь есть
	if !active {
		u.log.Infow("user not active")
		return "", errors.New("can't login")
	}

	token, _ := generateToken()

	return token, nil
}

func generateToken() (string, error) {
	return "sercretToken1234", nil
}

func generateHashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashBytes := hasher.Sum(nil)

	return hex.EncodeToString(hashBytes)
}
