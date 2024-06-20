package routes

import (
	database "capstone-project/database"
	"capstone-project/handler"
	"capstone-project/repository"
	"capstone-project/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	handler    handler.UserHandler
	otpHandler handler.OTPHandler
}

func NewUserRouter(handler handler.UserHandler, otpHandler handler.OTPHandler) *UserRouter {
	return &UserRouter{handler: handler, otpHandler: otpHandler}
}

func SetupUserRouter(router *gin.Engine, db *database.Database, redis *database.Redis) *UserRouter {
	userRepo := repository.NewUserRepository(db)
	sessionRepo := repository.NewSessionRepository(db)
	otpRepo := repository.NewOTPRepository(redis)

	userService := service.NewUserService(userRepo)
	_ = service.NewSessionService(sessionRepo)
	otpService := service.NewOTPService(otpRepo)

	userHandler := handler.NewUserHandler(userService)
	otpHandler := handler.NewOTPHandler(otpService, userService)

	userRouter := NewUserRouter(userHandler, otpHandler)

	version := router.Group("/api/v1")
	version.GET("/", userRouter.handler.GetServer)
	version.POST("/register", userRouter.handler.Register)
	version.POST("/login", userRouter.handler.Login)
	version.PATCH("/reset", userRouter.handler.ResetPassword)
	version.DELETE("/remove/:id", userRouter.handler.RemoveUser)

	otp := version.Group("/otp")
	otp.POST("/send", userRouter.otpHandler.SendOTP)
	otp.POST("/verify/:id", userRouter.otpHandler.VerifyOTP)

	return userRouter
}
