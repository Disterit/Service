package api

import (
	"Service/grpc/api/pb"
	"Service/grpc/internal/models"
	"context"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
}

func (s *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	var user models.Users
	user.Username = req.GetUsername()
	user.Password = req.GetPassword()

	return &pb.RegisterResponse{Message: "User registered successfully"}, nil
}
