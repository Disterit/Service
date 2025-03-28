package service

import (
	"Service/grpc/internal/models"
	"context"
)

type User interface {
	Register(ctx context.Context, user models.Users) error
}

type Server struct {
	User User
}

func NewServer(user User) *Server {
	return &Server{User: user}
}
