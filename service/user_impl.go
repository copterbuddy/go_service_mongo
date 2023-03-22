package service

import (
	"context"
	"main/repository"

	"github.com/google/uuid"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userService{repo: repo}
}

func (s userService) CreateUser(ctx context.Context, user repository.User) error {
	newUser := user
	newUser.ID = uuid.New().String()
	return s.repo.CreateUser(ctx, newUser)
}
