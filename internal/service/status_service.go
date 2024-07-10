package service

import (
	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"
	"github.com/andrew-nino/vtx_algorithms_synchronization/internal/repository/postgresdb"
)

type AlgorithmStatusService struct {
	repo postgresdb.AlgorithmStatusPostgres
}

func NewAlgorithmStatusService(repo postgresdb.AlgorithmStatusPostgres) *AlgorithmStatusService {
	return &AlgorithmStatusService{repo: repo}
}

func (s *AlgorithmStatusService) UpdateStatus(status entity.AlgorithmStatus) error {
	return s.repo.UpdateStatus(status)
}
