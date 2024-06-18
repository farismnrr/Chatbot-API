package routes

import (
	database "capstone-project/database"
	"capstone-project/handler"
	"capstone-project/repository"
	"capstone-project/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	handler handler.UserHandler
}

func NewUserRouter(handler handler.UserHandler) *UserRouter {
	return &UserRouter{handler: handler}
}

func SetupUserRouter(router *gin.Engine, db *database.Database) *UserRouter {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRouter := NewUserRouter(userHandler)

	version := router.Group("/api/v1")
	version.GET("/", userRouter.handler.GetServer)
	version.POST("/register", userRouter.handler.Register)

	return userRouter
}
