package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"main/handler"
	"main/repository"
	"main/service"
)

func main() {

	db := InitMongoDb()

	// Initialize Repositories
	userRepo := repository.NewUserMongoRepository(db)

	// Initialize Services
	userService := service.NewUserService(userRepo)

	// Initialize Handler
	userHandler := handler.NewUserHandler(userService)

	// Initialize Router
	router := gin.Default()

	router.POST("/users", userHandler.CreateUser)

	// Start Server
	if err := router.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

func InitMongoDb() *mongo.Database {
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:123456@localhost:27017/"))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {

		fmt.Println("====================")
		fmt.Println("cannot connect to mongodb!")
		fmt.Println("====================")

		log.Fatal(err)
	}

	fmt.Println("====================")
	fmt.Println("connected to mongodb!")
	fmt.Println("====================")

	return client.Database("mydb")
}
