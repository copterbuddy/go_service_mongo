package repository

import (
	"context"
)

type User struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"email" bson:"email"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user User) error
}
