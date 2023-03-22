package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, services ...interface{}) {

	var userHandler UserHandler

	if services == nil {
		panic("services nil")
	}

	for _, service := range services {
		switch item := service.(type) {
		case UserHandler:
			userHandler = item
		}
	}

	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, fmt.Sprintf("app running in version : beta"))
	})

	router1 := r.Group("")
	{
		router1.POST("/user", userHandler.CreateUser)
	}

}
