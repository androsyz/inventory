package transaction

import (
	"database/sql"

	"github.com/androsyz/inventory/config"
)

type Usecase struct {
	cfg             *config.Config
	sqlConn         *sql.DB
	repoTransaction repoTransactionInterface
	repoProduct     repoProductInterface
}

func New(cfg *config.Config, sqlConn *sql.DB, repoTransaction repoTransactionInterface, repoProduct repoProductInterface) *Usecase {
	return &Usecase{
		cfg:             cfg,
		sqlConn:         sqlConn,
		repoTransaction: repoTransaction,
		repoProduct:     repoProduct,
	}
}
