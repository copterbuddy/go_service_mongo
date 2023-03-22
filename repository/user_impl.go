package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	collection *mongo.Collection
}

func NewUserMongoRepository(db *mongo.Database) UserMongoRepository {
	return UserMongoRepository{
		collection: db.Collection("users"),
	}
}

func (r UserMongoRepository) CreateUser(ctx context.Context, user User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}
