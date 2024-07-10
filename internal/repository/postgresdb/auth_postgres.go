package postgresdb

import (
	"fmt"
	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// We create a new user in the database and return his ID or the error [ErrNoRows] if it does not work.
func (r *AuthPostgres) CreateManager(mng entity.Manager) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, managername, password_hash) values ($1, $2, $3) RETURNING id", managerTable)

	row := r.db.QueryRow(query, mng.Name, mng.Managername, mng.Password)
	if err := row.Scan(&id); err != nil {
		log.Debugf("repository.CreateManager - db.QueryRow : %v", err)
		return 0, err
	}

	return id, nil
}

// We make a request to the database about the user. An error is returned if the result set is empty.
func (r *AuthPostgres) GetManager(managerName, password string) (int, error) {
	var userID int
	query := fmt.Sprintf("SELECT id FROM %s WHERE managername=$1 AND password_hash=$2", managerTable)
	err := r.db.Get(&userID, query, managerName, password)
	if err != nil {
		log.Debugf("repository.CreateManager - db.Get : %v", err)
		return 0, err
	}
	return userID, err
}
