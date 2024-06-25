package handler

import (
	"capstone-project/model"
	"capstone-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type conversationHandler struct {
	userService         service.UserService
	sessionService      service.SessionService
	conversationService service.ConversationService
}

type ConversationHandler interface {
	GetConversation(c *gin.Context)
	DeleteConversation(c *gin.Context)
	DeleteAllConversation(c *gin.Context)
}

func NewConversationHandler(userService service.UserService, sessionService service.SessionService, conversationService service.ConversationService) *conversationHandler {
	return &conversationHandler{userService: userService, sessionService: sessionService, conversationService: conversationService}
}

func (h *conversationHandler) GetConversation(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	err = h.userService.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NewErrorResponse(http.StatusNotFound, "User not found"))
		return
	}

	conversations, err := h.conversationService.GetAllConversations(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, conversations)
}

func (h *conversationHandler) DeleteConversation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	_, err = h.conversationService.GetConversationById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NewErrorResponse(http.StatusNotFound, "Conversation not found"))
		return
	}

	err = h.conversationService.DeleteConversation(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse(http.StatusOK, "Conversation deleted"))
}

func (h *conversationHandler) DeleteAllConversation(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	err = h.userService.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NewErrorResponse(http.StatusNotFound, "User not found"))
		return
	}

	err = h.conversationService.DeleteAllConversation(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse(http.StatusOK, "All conversation deleted"))
}
