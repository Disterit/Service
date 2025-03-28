package api

import (
	"Service/grpc/api/pb"
	"Service/grpc/internal/models"
	"Service/grpc/internal/service"
	"context"
	"go.uber.org/zap"
)

type AuthHandler struct {
	user service.User
	log  *zap.SugaredLogger
	pb.UnimplementedAuthServiceServer
}

func NewAuthHandler(user service.User, log *zap.SugaredLogger) *AuthHandler {
	return &AuthHandler{
		user: user,
		log:  log,
	}
}

func (s *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	var user models.Users
	user.Username = req.GetUsername()
	user.Password = req.GetPassword()
	user.Email = req.GetEmail()

	err := s.user.Register(ctx, user)
	if err != nil {
		s.log.Errorw("error to create user in service", "error", err)
		return nil, err
	}

	s.log.Infow("register user in service", "username", user.Username)

	return &pb.RegisterResponse{Message: "User registered successfully"}, nil
}

func (s *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.Users
	user.Username = req.GetUsername()
	user.Password = req.GetPassword()

	token, err := s.user.Login(ctx, user)
	if err != nil {
		s.log.Errorw("error to login in service", "error", err)
		return nil, err
	}

	s.log.Infow("login in service", "username", user.Username)

	return &pb.LoginResponse{Token: token}, nil
}
