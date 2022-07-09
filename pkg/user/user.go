package user

import (
	"context"
	"msqrd/user-service/pkg/api/userapi"
)

func (s *GRPCServer) CreateUser(ctx context.Context, r *userapi.CreateUserRequest) (*userapi.CreateUserResponse, error) {
	return nil, nil
}
