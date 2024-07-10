package postgresdb

import (
	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateManager(user entity.Manager) (int, error)
	GetManager(username, password string) (int, error)
}

type ClientPostgres interface {
	AddClient(client entity.Client) (int, error)
	UpdateClient(client entity.Client) (int, error)
	DeleteClient(id int) error
}

type AlgorithmStatusPostgres interface {
	UpdateStatus(status entity.AlgorithmStatus) error
}

type PG_Repository struct {
	Authorization
	ClientPostgres
	AlgorithmStatusPostgres
}

func NewPGRepository(db *sqlx.DB) *PG_Repository {
	return &PG_Repository{
		Authorization:           NewAuthPostgres(db),
		ClientPostgres:          NewClientToPostgres(db),
		AlgorithmStatusPostgres: NewAlgorithmStatusToPostgres(db),
	}
}
