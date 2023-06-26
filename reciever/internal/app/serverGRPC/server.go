package serverGRPC

import (
	"context"
	"fmt"
	"github.com/sinyavcev/proto/pb"
	"reciever/common/logger"
	"reciever/internal/repository"
)

type Server struct {
	Logger                            *logger.Logger
	Repos                             *repository.Repository
	pb.UnimplementedUserServiceServer // встраивание не реализованного интерфейса
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	err := s.Repos.UserRepository.CreateUser(ctx, *req)
	if err != nil {
		return nil, fmt.Errorf("CreateUser", err)
	}
	result := "Пользователь " + req.Name + " добавлен"

	s.Logger.Info(result)

	return &pb.User{
		Name:  req.Name,
		Phone: req.Phone,
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	userName, err := s.Repos.UserRepository.UpdateUser(ctx, *req)
	if err != nil {
		return nil, fmt.Errorf("update User %w", err)
	}
	result := "Пользователь " + userName + " обновлен на " + req.Name

	s.Logger.Info(result)

	return &pb.User{
		Name:  req.Name,
		Phone: req.Phone,
	}, nil
}
