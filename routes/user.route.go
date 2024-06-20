package routes

import (
	database "capstone-project/database"
	"capstone-project/handler"
	"capstone-project/repository"
	"capstone-project/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userHandler handler.UserHandler
	otpHandler  handler.OTPHandler
}

func NewUserRouter(userHandler handler.UserHandler, otpHandler handler.OTPHandler) *UserRouter {
	return &UserRouter{userHandler: userHandler, otpHandler: otpHandler}
}

func SetupUserRouter(router *gin.Engine, db *database.Database, redis *database.Redis) *UserRouter {
	userRepo := repository.NewUserRepository(db)
	sessionRepo := repository.NewSessionRepository(redis)
	otpRepo := repository.NewOTPRepository(redis)

	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(sessionRepo)
	otpService := service.NewOTPService(otpRepo)

	userHandler := handler.NewUserHandler(userService, sessionService)
	otpHandler := handler.NewOTPHandler(otpService, userService)

	userRouter := NewUserRouter(userHandler, otpHandler)

	version := router.Group("/api/v1")
	version.GET("/", userRouter.userHandler.GetServer)
	version.POST("/register", userRouter.userHandler.Register)
	version.POST("/login", userRouter.userHandler.Login)
	version.PATCH("/reset", userRouter.userHandler.ResetPassword)

	// version.Use(middleware.AuthMiddleware())
	version.DELETE("/remove/:id", userRouter.userHandler.RemoveUser)
	version.POST("/request", userRouter.userHandler.RequestAPI)

	otp := version.Group("/otp")
	otp.POST("/send", userRouter.otpHandler.SendOTP)
	otp.POST("/verify/:id", userRouter.otpHandler.VerifyOTP)

	return userRouter
}
