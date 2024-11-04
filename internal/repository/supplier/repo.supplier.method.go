package supplier

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/androsyz/inventory/internal/model"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
)

func (r *Repository) CreateSupplier(ctx context.Context, supplier *model.Supplier) (int, error) {
	var currentInsertedID int

	values := []interface{}{
		supplier.Name,
		supplier.LeadtimeMax,
		supplier.LeadtimeAvg,
	}

	err := r.sqlConn.QueryRowContext(ctx, sqlInsertSupplier, values...).Scan(&currentInsertedID)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when create supplier, got %v", err))
		return 0, err
	}

	return currentInsertedID, nil
}

func (r *Repository) GetSuppliers(ctx context.Context) ([]*model.Supplier, error) {
	var suppliers []*model.Supplier

	rows, err := r.sqlConn.QueryContext(ctx, sqlGetSuppliers)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when get suppliers, got %v", err))
		return nil, err
	}

	for rows.Next() {
		temp := model.Supplier{}
		err := rows.Scan(&temp.ID, &temp.Name, &temp.LeadtimeMax, &temp.LeadtimeAvg)
		if err != nil {
			zlog.Error(ctx, nil, fmt.Sprintf("error when scan supplier, got %v", err))
		}

		suppliers = append(suppliers, &temp)
	}

	return suppliers, nil
}

func (r *Repository) GetSupplierByID(ctx context.Context, id int) (*model.Supplier, error) {
	supplier := &model.Supplier{}

	row := r.sqlConn.QueryRowContext(ctx, sqlGetSupplierByID, id)

	err := row.Scan(&supplier.ID, &supplier.Name, &supplier.LeadtimeMax, &supplier.LeadtimeAvg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		zlog.Error(ctx, nil, fmt.Sprintf("error when scan: %v", err))
		return nil, err
	}

	return supplier, nil
}
