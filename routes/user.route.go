package routes

import (
	database "capstone-project/database"
	"capstone-project/handler"
	"capstone-project/middleware"
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
	sessionRepo := repository.NewSessionRepository(db)
	userService := service.NewUserService(userRepo, sessionRepo)
	userHandler := handler.NewUserHandler(userService)
	userRouter := NewUserRouter(userHandler)

	version := router.Group("/api/v1")
	version.GET("/", userRouter.handler.GetServer)
	version.POST("/register", userRouter.handler.Register)
	version.POST("/login", userRouter.handler.Login)
	version.Use(middleware.Auth())
	version.DELETE("/remove/:id", userRouter.handler.RemoveUser)

	return userRouter
}
