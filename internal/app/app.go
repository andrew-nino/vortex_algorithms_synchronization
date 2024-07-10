package app

import (
	"vtx_algorithms_synchronization/config"
	"vtx_algorithms_synchronization/pkg/postgres"

	log "github.com/sirupsen/logrus"
)

// Initialization and start of critical components.
func Run(configPath string) {

	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	SetLogrus(cfg.Log.Level)

	// Repositories
	log.Info("Initializing postgres...")
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	// Migrates running
	log.Info("Migrates running...")
	m := NewMigration(cfg)
	err = m.Steps(1)
	if err != nil {
		log.Fatalf("failed to migrate db: %s", err.Error())
	}
}
