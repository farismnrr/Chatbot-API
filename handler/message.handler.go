package handler

import (
	"capstone-project/model"
	"capstone-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type messageHandler struct {
	messageService      service.MessageService
	conversationService service.ConversationService
	sessionService      service.SessionService
}

type MessageHandler interface {
	GetMessage(c *gin.Context)
	DeleteMessage(c *gin.Context)
}

func NewMessageHandler(messageService service.MessageService, sessionService service.SessionService, conversationService service.ConversationService) *messageHandler {
	return &messageHandler{messageService: messageService, sessionService: sessionService, conversationService: conversationService}
}

func (h *messageHandler) GetMessage(c *gin.Context) {
	convID, err := strconv.Atoi(c.Param("conversation_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid conversation ID"))
		return
	}

	_, err = h.conversationService.GetConversationById(convID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NewErrorResponse(http.StatusNotFound, "Conversation not found"))
		return
	}

	messages, err := h.messageService.GetMessage(convID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (h *messageHandler) DeleteMessage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Invalid conversation ID"))
		return
	}

	_, err = h.messageService.GetMessageById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.NewErrorResponse(http.StatusNotFound, "Message not found"))
		return
	}

	err = h.messageService.DeleteMessage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse(http.StatusOK, "Message deleted"))
}
