package service

import (
	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"
	"github.com/andrew-nino/vtx_algorithms_synchronization/internal/deployment"
	"github.com/andrew-nino/vtx_algorithms_synchronization/internal/repository/postgresdb"

	log "github.com/sirupsen/logrus"
)

var mapStatuses = make(map[int64][]string)

type AlgorithmStatusService struct {
	repo postgresdb.AlgorithmStatusPostgres
}

func NewAlgorithmStatusService(repo postgresdb.AlgorithmStatusPostgres) *AlgorithmStatusService {
	return &AlgorithmStatusService{repo: repo}
}

func (s *AlgorithmStatusService) UpdateStatus(status entity.AlgorithmStatus) error {
	return s.repo.UpdateStatus(status)
}

func (s *AlgorithmStatusService) CheckAlgorithmStatus(deploy *deployment.Deploy) {

	statusClients, err := s.repo.CheckAlgorithmStatus()
	if err != nil {
		log.Errorf("StatusService.CheckAlgorithmStatus - s.repo.CheckAlgorithmStatus: %v", err)
	}

	for _, client := range statusClients {

		if client.VWAP {
			CheckAndStartDeployment(client.ClientID, "VWAP", mapStatuses, deploy)
		} else {
			CheckAndStopDeployment(client.ClientID, "VWAP", mapStatuses, deploy)
		}

		if client.TWAP {
			CheckAndStartDeployment(client.ClientID, "TWAP", mapStatuses, deploy)
		} else {
			CheckAndStopDeployment(client.ClientID, "TWAP", mapStatuses, deploy)
		}

		if client.HFT {
			CheckAndStartDeployment(client.ClientID, "HFT", mapStatuses, deploy)
		} else {
			CheckAndStopDeployment(client.ClientID, "HFT", mapStatuses, deploy)
		}

	}

}

func CheckAndStartDeployment(clientID int64, statusClients string, mapStatuses map[int64][]string, deploy *deployment.Deploy) {

	if statuses, ok := mapStatuses[clientID]; ok {
		for _, s := range statuses {
			if s == statusClients {
				return
			}
		}
		deploy.CreatePod(clientID, statusClients)
		mapStatuses[clientID] = append(mapStatuses[clientID], statusClients)

	} else {
		deploy.CreatePod(clientID, statusClients)
		mapStatuses[clientID] = append(mapStatuses[clientID], statusClients)
	}
}

func CheckAndStopDeployment(clientID int64, statusClients string, mapStatuses map[int64][]string, deploy *deployment.Deploy) {

	if statuses, ok := mapStatuses[clientID]; ok {
		for i := 0; i < len(statuses); i++ {
			if statuses[i] == statusClients {
				deploy.DeletePod(clientID, statusClients)
				statuses = append(statuses[:i], statuses[i+1:]...)
				mapStatuses[clientID] = statuses
				return
			}
		}
	}
}
