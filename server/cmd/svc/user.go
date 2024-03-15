package svc

import (
	"context"
	"server/proto/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	user.UnimplementedUserServiceServer
}

func (s *UserServiceServer) CreateUser(ctx context.Context,
	req *user.CreateUserRequest) (res *user.CreateUserResponse, err error) {

	if req.Name == "" {
		err = status.New(codes.InvalidArgument, "Name cannot be empty").Err()
		return nil, err
	}

	res = &user.CreateUserResponse{
		Name: req.Name,
	}

	return res, nil
}
