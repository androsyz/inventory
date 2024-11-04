package transaction

import (
	"context"
	"database/sql"

	"github.com/androsyz/inventory/internal/model"
)

type repoTransactionInterface interface {
	CreateTransaction(ctx context.Context, tx *sql.Tx, transaction *model.Transaction) (int, error)
}

type repoProductInterface interface {
	GetProductByID(ctx context.Context, tx *sql.Tx, id int) (*model.Product, error)
	UpdateProductStock(ctx context.Context, tx *sql.Tx, stock int, id int) error
}
