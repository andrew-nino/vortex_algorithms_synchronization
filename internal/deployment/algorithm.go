package deployment

import "time"

type Algorithm struct {
	Name      string
	CreatedAt time.Time
}

func NewAlgorithm(name string) Algorithm {
	return Algorithm{Name: name, CreatedAt: time.Now()}
}
