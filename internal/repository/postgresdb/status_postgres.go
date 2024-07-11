package postgresdb

import (
	"fmt"

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
	tx, err := as.db.Begin()
	if err != nil {
		return err
	}

	var checkID int
	query := "UPDATE algorithm_status SET vwap=$1, twap=$2, hft=$3 WHERE client_id=$4 RETURNING id"
	row := tx.QueryRow(query, status.VWAP, status.TWAP, status.HFT, status.ClientID)
	err = row.Scan(&checkID)
	if err != nil {
		tx.Rollback()
		log.Debugf("repository.UpdateStatus - row.Scan: %v", err)
		return err
	}
	return tx.Commit()
}

func (as *AlgorithmStatusToPostgres) CheckAlgorithmStatus() ([]entity.AlgorithmStatus, error) {

	clientsStatuses := []entity.AlgorithmStatus{}

	query := fmt.Sprintf("SELECT client_id, vwap, twap, hft FROM %s ", statusTable)
	err := as.db.Select(&clientsStatuses, query)
	if err != nil {
		log.Debugf("repository.CheckAlgorithmStatus - db.Select: %v", err)
		return nil, err
	}

	return clientsStatuses, nil
}
