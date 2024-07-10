package service

import (
	"github.com/andrew-nino/vtx_algorithms_synchronization/config"
	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"

	postgres "github.com/andrew-nino/vtx_algorithms_synchronization/internal/repository/postgresdb"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateManager(user entity.Manager) (int, error)
	SignIn(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(reposPG *postgres.PG_Repository, cfg *config.Config) *Service {
	return &Service{
		Authorization: NewAuthService(reposPG.Authorization, cfg),
	}
}
