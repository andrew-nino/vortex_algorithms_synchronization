package postgresdb

import (
	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type AlgorithmStatusToPostgres struct {
	db *sqlx.DB
}

func NewAlgorithmStatusToPostgres(db *sqlx.DB) *AlgorithmStatusToPostgres {
	return &AlgorithmStatusToPostgres{db: db}
}

func (as *AlgorithmStatusToPostgres) UpdateStatus(status entity.AlgorithmStatus) error {

	var checkID int
	query := "UPDATE algorithm_status SET vwap=$1, twap=$2, hft=$3 WHERE client_id=$4 RETURNING id"
	row := as.db.QueryRow(query, status.VWAP, status.TWAP, status.HFT, status.ClientID)
	err := row.Scan(&checkID)
	if err != nil {
		log.Debugf("repository.UpdateStatus - row.Scan: %v", err)
		return err
	}
	return nil
}
