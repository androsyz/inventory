package supplier

import (
	"database/sql"

	"github.com/androsyz/inventory/config"
)

type Repository struct {
	cfg     *config.Config
	sqlConn *sql.DB
}

func New(cfg *config.Config, sqlConn *sql.DB) *Repository {
	repo := &Repository{
		cfg:     cfg,
		sqlConn: sqlConn,
	}

	return repo
}
