package postgresdb

import (
	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateManager(user entity.Manager) (int, error)
	GetManager(username, password string) (int, error)
}

type PG_Repository struct {
	Authorization
}

func NewPGRepository(db *sqlx.DB) *PG_Repository {
	return &PG_Repository{
		Authorization: NewAuthPostgres(db),
	}
}
