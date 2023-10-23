package service

import (
	"context"
	"github.com/6a6ydoping/deeznuts_push_service/internal/entity"
)

type UserService interface {
}

type MessageService interface {
	SendMessage(ctx context.Context, payload *entity.PushPayload) (string, error)
}
