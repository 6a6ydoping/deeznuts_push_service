package service

import (
	"github.com/6a6ydoping/deeznuts_push_service/internal/config"
	"github.com/6a6ydoping/deeznuts_push_service/internal/repository"
	"net/http"
)

type Manager struct {
	Repository repository.Repository
	Config     *config.Config
	Client     *http.Client
}

func New(repository repository.Repository, config *config.Config, client *http.Client) *Manager {
	return &Manager{
		Repository: repository,
		Config:     config,
		Client:     client,
	}
}
