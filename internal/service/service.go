package service

import (
	"github.com/andrew-nino/vtx_algorithms_synchronization/config"
	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"

	postgres "github.com/andrew-nino/vtx_algorithms_synchronization/internal/repository/postgresdb"
)

type Authorization interface {
	CreateManager(user entity.Manager) (int, error)
	SignIn(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Client interface {
	AddClient(client entity.Client) (int, error)
	UpdateClient(client entity.Client) (int, error)
	DeleteClient(id int) error
}

type AlgorithmStatus interface {
	UpdateStatus(status entity.AlgorithmStatus) error
}

type Service struct {
	Authorization
	Client
	AlgorithmStatus
}

func NewService(reposPG *postgres.PG_Repository, cfg *config.Config) *Service {
	return &Service{
		Authorization:   NewAuthService(reposPG.Authorization, cfg),
		Client:          NewClientService(reposPG.ClientPostgres),
		AlgorithmStatus: NewAlgorithmStatusService(reposPG.AlgorithmStatusPostgres),
	}
}
