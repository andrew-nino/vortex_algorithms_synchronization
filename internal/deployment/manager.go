package deployment

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type DeployManager struct {
	Pods []Pod
}

func NewDeployManager() *DeployManager {
	return &DeployManager{Pods: make([]Pod, 0)}
}

func (dm *DeployManager) CreatePod(clientID int64, algorithmStr string) error {
	if len(algorithmStr) == 0 {
		return fmt.Errorf("algorithm name must be provided")
	}
	algorithm := NewAlgorithm(algorithmStr)
	newPod := NewPod(clientID, algorithm)

	dm.Pods = append(dm.Pods, newPod)
	log.Printf("A pod for the algorithm %s for the client %d has been created", algorithm.Name, newPod.ClientID)
	return nil
}

func (dm *DeployManager) DeletePod(clientID int64, algorithmStr string) error {

	for i, pod := range dm.Pods {
		if pod.ClientID == clientID && pod.Algorithm.Name == algorithmStr {
			dm.Pods = append(dm.Pods[:i], dm.Pods[i+1:]...)
			log.Printf("Deleted pod with ClientID = %d and algorithm name = %s", clientID, algorithmStr)
			return nil
		}
	}
	fmt.Printf("Pod with ClientID = %d not found\n", clientID)
	return nil
}

func (dm *DeployManager) GetPodList() ([]string, error) {
	podNames := make([]string, len(dm.Pods))
	for i, pod := range dm.Pods {
		podNames[i] = fmt.Sprintf("ClientID = %d  algorithm = %s",pod.ClientID, pod.Algorithm.Name)
	}
	return podNames, nil
}
