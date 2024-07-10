package postgres

import (
	"fmt"
	"github.com/andrew-nino/vtx_algorithms_synchronization/config"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

type Postgres struct {
	connAttempts int
	connTimeout  time.Duration
}

// Causes the database to open and checks the connection. If the connection is established, returns a pointer to the database.
// Returns an error if the database has not opened or there is no connection.
func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {

	pg := &Postgres{
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PG.Host, cfg.PG.Port, cfg.PG.Username, cfg.PG.DBName, cfg.PG.Password, cfg.PG.SSLMode))
	if err != nil {
		return nil, err
	}

	for pg.connAttempts > 0 {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)
		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}

	return db, err
}
