package transaction

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/androsyz/inventory/internal/model"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
)

func (r *Repository) CreateTransaction(ctx context.Context, tx *sql.Tx, transaction *model.Transaction) (int, error) {
	var currentInsertedID int

	values := []interface{}{
		transaction.UserID,
		transaction.ProductID,
		transaction.TotalPrice,
		transaction.Quantity,
	}

	err := tx.QueryRowContext(ctx, sqlInsertTransaction, values...).Scan(&currentInsertedID)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when create transaction, got %v", err))
		return 0, err
	}

	return currentInsertedID, nil
}
