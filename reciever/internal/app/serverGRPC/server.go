package serverGRPC

import (
	"context"
	"reciever/pb"
)

type Server struct {
	pb.UnimplementedUserServiceServer // встраивание не реализованного интерфейса
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	return nil, nil
}
