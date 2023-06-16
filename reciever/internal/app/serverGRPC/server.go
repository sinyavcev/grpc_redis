package serverGRPC

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"reciever/internal/repository"
	"reciever/pb"
)

type Server struct {
	Repos                             repository.Repository
	pb.UnimplementedUserServiceServer // встраивание не реализованного интерфейса
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*wrappers.StringValue, error) {
	err := s.Repos.UserRepository.CreateUser(ctx, *req)
	if err != nil {
		return &wrappers.StringValue{Value: "Пользователь не добавлен"}, fmt.Errorf("CreateUser", err)
	}
	result := "Пользователь" + req.Name + " добавлен"
	fmt.Println(result)
	return &wrappers.StringValue{Value: result}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*wrappers.StringValue, error) {
	userName, err := s.Repos.UserRepository.UpdateUser(ctx, *req)
	if err != nil {
		return &wrappers.StringValue{Value: "Пользователь не обновлен"}, fmt.Errorf("CreateUser", err)
	}
	result := "Пользователь " + userName + " обновлен на " + req.Name
	fmt.Println(result)
	return &wrappers.StringValue{Value: result}, nil
}
