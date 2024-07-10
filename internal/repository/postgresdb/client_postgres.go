package postgresdb

import (
	"fmt"
	"strings"

	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type ClientToPostgres struct {
	db *sqlx.DB
}

func NewClientToPostgres(db *sqlx.DB) *ClientToPostgres {
	return &ClientToPostgres{db: db}
}

func (c *ClientToPostgres) AddClient(add entity.Client) (int, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return 0, err
	}
	var clientID int
	queryToClient := fmt.Sprintf(`INSERT INTO %s (client_id, client_name, version, image, cpu, memory, priority, needRestart, spawned_at) 
									values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`, clientTable)
	rowClient := tx.QueryRow(queryToClient, add.ID, add.ClientName, add.Version, add.Image, add.CPU, add.Memory, add.Priority, add.NeedRestart, add.SpawnedAt)
	err = rowClient.Scan(&clientID)
	if err != nil {
		log.Debugf("repository.AddClient - rowClient.Scan : %v", err)
		tx.Rollback()
		return 0, err
	}
	queryToStatus := fmt.Sprintf(`INSERT INTO %s (client_id) values ($1) RETURNING id`, statusTable)
	_, err = tx.Exec(queryToStatus, add.ID)
	if err != nil {
		tx.Rollback()
		log.Debugf("repository.AddClient - tx.Exec : %v", err)
		return 0, err
	}

	return clientID, tx.Commit()
}

func (c *ClientToPostgres) UpdateClient(client entity.Client) (int, error) {
	var id int
	setValues := make([]string, 0)

	if client.ID == 0 {
		return 0, fmt.Errorf("client id must be provided")
	}

	if client.ClientName == "" {
		return 0, fmt.Errorf("client name must be provided")
	} else {
		setValues = append(setValues, fmt.Sprintf("client_name='%s'", client.ClientName))
	}

	if client.Version != 0 {
		setValues = append(setValues, fmt.Sprintf("version='%d'", client.Version))
	}

	if client.Image != "" {
		setValues = append(setValues, fmt.Sprintf("image='%s'", client.Image))
	}
	if client.CPU != "" {
		setValues = append(setValues, fmt.Sprintf("cpu='%s'", client.CPU))
	}
	if client.Memory != "" {
		setValues = append(setValues, fmt.Sprintf("memory='%s'", client.Memory))
	}
	if client.Priority != 0.0 {
		setValues = append(setValues, fmt.Sprintf("priority='%f'", client.Priority))
	}
	if client.NeedRestart {
		setValues = append(setValues, fmt.Sprintf("needRestart='%v'", client.NeedRestart))
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s  WHERE client_id = $1 RETURNING id`, clientTable, setQuery)

	row := c.db.QueryRow(query, client.ID)
	err := row.Scan(&id)
	if err != nil {
		log.Debugf("repository.UpdateClient - row.Scan : %v", err)
		return 0, err
	}

	return id, nil
}

func (c *ClientToPostgres) DeleteClient(clientId int) error {
	var checkID int
	query := fmt.Sprintf(`DELETE FROM %s WHERE client_id = $1 RETURNING id`, clientTable)
	row := c.db.QueryRow(query, clientId)
	err := row.Scan(&checkID)
	if err != nil {
		log.Debugf("repository.DeleteClient - row.Scan : %v", err)
		return err
	}
	return nil
}
