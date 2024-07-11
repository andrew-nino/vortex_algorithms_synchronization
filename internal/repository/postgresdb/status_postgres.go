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

// The status is updated if there is a match on client_id.
// The work is wrapped in a transaction, since the data from the table is used by the main method of checking statuses for changes.
func (as *AlgorithmStatusToPostgres) UpdateStatus(status entity.AlgorithmStatus) error {
	tx, err := as.db.Begin()
	if err != nil {
		return err
	}

	var checkID int
	query := fmt.Sprintf("UPDATE %s SET vwap=$1, twap=$2, hft=$3 WHERE client_id=$4 RETURNING id", statusTable)
	row := tx.QueryRow(query, status.VWAP, status.TWAP, status.HFT, status.ClientID)
	err = row.Scan(&checkID)
	if err != nil {
		tx.Rollback()
		log.Debugf("repository.UpdateStatus - row.Scan: %v", err)
		return err
	}
	return tx.Commit()
}

// We select data based on the status of the client’s algorithms.
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
