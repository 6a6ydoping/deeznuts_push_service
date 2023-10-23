package handler

import "github.com/6a6ydoping/deeznuts_push_service/internal/service"

type Handler struct {
	userService    service.UserService
	messageService service.MessageService
}

func New(us service.UserService, ms service.MessageService) *Handler {
	return &Handler{
		userService:    us,
		messageService: ms,
	}
}
