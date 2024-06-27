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
	userHandler         handler.UserHandler
	otpHandler          handler.OTPHandler
	conversationHandler handler.ConversationHandler
	messageHandler      handler.MessageHandler
}

func NewUserRouter(userHandler handler.UserHandler, otpHandler handler.OTPHandler, conversationHandler handler.ConversationHandler, messageHandler handler.MessageHandler) *UserRouter {
	return &UserRouter{userHandler: userHandler, otpHandler: otpHandler, conversationHandler: conversationHandler, messageHandler: messageHandler}
}

func SetupUserRouter(router *gin.Engine, db *database.Database, redis *database.Redis) *UserRouter {
	userRepo := repository.NewUserRepository(db)
	sessionRepo := repository.NewSessionRepository(redis)
	otpRepo := repository.NewOTPRepository(redis)
	conversationRepo := repository.NewConversationRepository(db)
	messageRepo := repository.NewMessageRepository(db)

	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(sessionRepo)
	otpService := service.NewOTPService(otpRepo)
	conversationService := service.NewConversationService(conversationRepo)
	messageService := service.NewMessageService(messageRepo)

	userHandler := handler.NewUserHandler(userService, sessionService, conversationService, messageService)
	conversationHandler := handler.NewConversationHandler(userService, sessionService, conversationService)
	messageHandler := handler.NewMessageHandler(messageService, sessionService, conversationService)
	otpHandler := handler.NewOTPHandler(otpService, userService)

	userRouter := NewUserRouter(userHandler, otpHandler, conversationHandler, messageHandler)

	router.GET("/", userRouter.userHandler.GetServer)
	version := router.Group("/api/v1")

	user := version.Group("/user")
	user.POST("/register", userRouter.userHandler.Register)
	user.POST("/login", userRouter.userHandler.Login)
	user.PATCH("/reset", userRouter.userHandler.ResetPassword)

	user.Use(middleware.AuthMiddleware())
	user.DELETE("/logout/:id", userRouter.userHandler.Logout)
	user.DELETE("/remove/:id", userRouter.userHandler.RemoveUser)
	user.POST("/conversation", userRouter.userHandler.CreateConversation)
	user.POST("/conversation/:conversation_id", userRouter.userHandler.CreateMessage)

	user.GET("/conversation/:user_id", userRouter.conversationHandler.GetConversation)
	user.DELETE("/conversation/:id", userRouter.conversationHandler.DeleteConversation)
	user.DELETE("/conversations/:user_id", userRouter.conversationHandler.DeleteAllConversation)

	user.GET("/message/:conversation_id", userRouter.messageHandler.GetMessage)
	user.DELETE("/message/:id", userRouter.messageHandler.DeleteMessage)

	otp := version.Group("/otp")
	otp.POST("/send", userRouter.otpHandler.SendOTP)
	otp.POST("/verify/:id", userRouter.otpHandler.VerifyOTP)

	return userRouter
}
