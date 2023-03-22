package service

import (
	"context"
	"main/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user repository.User) error
}
