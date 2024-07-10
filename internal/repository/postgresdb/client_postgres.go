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

func NewClientPostgres(db *sqlx.DB) *ClientToPostgres {
	return &ClientToPostgres{db: db}
}

func (c *ClientToPostgres) AddClient(add entity.Client) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (client_id, client_name, version, image, cpu, memory, priority, needRestart, spawned_at) 
									values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`, clientTable)
	row := c.db.QueryRow(query, add.ID, add.ClientName, add.Version, add.Image, add.CPU, add.Memory, add.Priority, add.NeedRestart, add.SpawnedAt)
	err := row.Scan(&id)
	if err != nil {
		log.Debugf("repository.AddClient - row.Scan : %v", err)
		return 0, err
	}
	return id, nil
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

	query := fmt.Sprintf(`DELETE FROM %s WHERE client_id = $1 `, clientTable)
	_, err := c.db.Exec(query, clientId)
	if err != nil {
		log.Debugf("repository.DeleteClient - db.Exec : %v", err)
		return err
	}
	return nil
}
