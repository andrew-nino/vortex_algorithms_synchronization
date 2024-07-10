package app

import (
	"log"
	"vtx_algorithms_synchronization/config"

	go_migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Returns a new Migrate instance
func NewMigration(c *config.Config) *go_migrate.Migrate {
	m, err := go_migrate.New(
		"file://schema",
		"postgres://"+c.PG.Username+":"+c.PG.Password+"@"+c.PG.Host+":"+c.PG.Port+"/"+c.PG.DBName+"?sslmode="+c.PG.SSLMode+"")

	if err != nil {
		log.Fatalf("failed to migrate db: %s", err.Error())
	}

	return m
}
