package deployment

type Deployer interface {
	CreatePod(clientID int64, name string) error
	DeletePod(clientID int64, name string) error
	GetPodList() ([]string, error)
}

type Deploy struct {
	Deployer
}

func NewDeploy() *Deploy {
	return &Deploy{Deployer: NewDeployManager()}
}
