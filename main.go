package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"email" bson:"email"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, user *User) error {
	newUser := user
	newUser.ID = uuid.New().String()
	return s.repo.CreateUser(ctx, newUser)
}

type UserMongoRepository struct {
	collection *mongo.Collection
}

func (r *UserMongoRepository) CreateUser(ctx context.Context, user *User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func NewUserMongoRepository(db *mongo.Database) *UserMongoRepository {
	return &UserMongoRepository{
		collection: db.Collection("users"),
	}
}

func main() {
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:123456@localhost:27017/"))
	if err != nil {
		panic(err)
	}

	db := client.Database("mydb")

	// Initialize Repositories
	userRepo := NewUserMongoRepository(db)

	// Initialize Services
	userSvc := &UserService{
		repo: userRepo,
	}

	// Initialize Router
	router := gin.Default()

	router.POST("/users", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := userSvc.CreateUser(c.Request.Context(), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, user)
	})

	// Start Server
	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
