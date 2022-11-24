package http

import (
	"vinbigdata/internal/delivery/http/controller"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ginApp *gin.Engine

	telecomController *controller.TelecomController
}

type HandlerOption func(*Handler)

func NewHandler(options ...HandlerOption) *Handler {
	handler := &Handler{}
	for _, option := range options {
		option(handler)
	}
	return handler
}

func WithGinEngine(r *gin.Engine) HandlerOption {
	return func(handler *Handler) {
		handler.ginApp = r
	}
}

func WithTelecomController(c *controller.TelecomController) HandlerOption {
	return func(handler *Handler) {
		handler.telecomController = c
	}
}
