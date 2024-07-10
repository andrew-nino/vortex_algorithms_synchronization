package service

import (
	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"
	"github.com/andrew-nino/vtx_algorithms_synchronization/internal/repository/postgresdb"
)

type ClientService struct {
	repo postgresdb.ClientPostgres
}

func NewClientService(repo postgresdb.ClientPostgres) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) AddClient(client entity.Client) (int, error) {
	return s.repo.AddClient(client)
}

func (s *ClientService) UpdateClient(client entity.Client) (int, error) {
	return s.repo.UpdateClient(client)
}

func (s *ClientService) DeleteClient(clientID int) error {
	return s.repo.DeleteClient(clientID)
}
