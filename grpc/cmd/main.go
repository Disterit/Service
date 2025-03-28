package main

import (
	"Service/grpc/api/pb"
	"Service/grpc/internal/api"
	"Service/grpc/internal/config"
	logging "Service/grpc/internal/logger"
	"Service/grpc/internal/repository"
	"Service/grpc/internal/service"
	"context"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// подгрузка env файла
	if err := godotenv.Load(config.EnvPath); err != nil {
		log.Fatal("Error loading .env file")
	}

	// считывание информации в config
	var cfg config.AppConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal("Error processing config", zap.Error(err))
	}

	// создаем свой логгер
	logger, err := logging.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal("Error init logger", zap.Error(err))
	}

	// соединение с БД
	pool, err := repository.Connection(context.Background(), cfg.PostgresDB)
	if err != nil {
		log.Fatal("Error connect to database", zap.Error(err))
	}

	// прооверка соединения с бд
	if err := repository.CheckConnection(pool, logger); err != nil {
		log.Fatalf("Connection check failed: %v", err)
	}

	// создание структуры пользователя для репозитория
	userRepository := repository.NewUserRepository(pool)
	repos := repository.NewRepository(userRepository)

	// создание структуры пользователя для сервиса
	userService := service.NewUserService(repos.User, logger)
	services := service.NewService(userService)

	// создание уровня handler
	handlers := api.NewAuthHandler(services.User, logger)

	server := grpc.NewServer()

	pb.RegisterAuthServiceServer(server, handlers)

	lis, err := net.Listen("tcp", cfg.Grpc.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting gRPC server on port 50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
