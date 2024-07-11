package deployment

type Pod struct {
	ClientID  int64
	Algorithm Algorithm
	IsRunning bool
}

func NewPod(clientID int64, alg Algorithm) Pod {
	return Pod{ClientID: clientID, Algorithm: alg, IsRunning: true}
}
