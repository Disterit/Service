package service

import (
	"Service/grpc/internal/models"
	"Service/grpc/internal/repository"
	"context"
	"errors"
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

	ok := u.repo.CheckExists(ctx, user)
	if !ok {
		u.log.Info("User already exists")
		return errors.New("user already created")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Errorw("Failed to hash password for registration", "error", err)
		return err
	}

	user.Password = string(hashPassword)

	err = u.repo.Register(ctx, user)
	if err != nil {
		u.log.Errorw("error to create user in db", "error", err)
		return err
	}

	return nil
}

func (u *userService) Login(ctx context.Context, user models.Users) (string, error) {

	userGet, err := u.repo.Login(ctx, user)
	if err != nil {
		u.log.Errorw("error to login", "error", err)
		return "", err
	}

	compareHashes := bcrypt.CompareHashAndPassword([]byte(userGet.Password), []byte(user.Password))
	if compareHashes != nil {
		u.log.Errorw("error to login, check password or name", "error", errors.New("invalid password"))
		return "", errors.New("invalid password")
	}

	// тут логика создание токена будет если пользователь есть
	if !userGet.Active {
		u.log.Infow("user not active")
		return "", errors.New("can't login")
	}

	token, _ := generateToken()

	return token, nil
}

func generateToken() (string, error) {
	return "sercretToken1234", nil
}
