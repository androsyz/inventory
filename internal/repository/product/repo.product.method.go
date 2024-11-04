package product

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/androsyz/inventory/internal/model"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
)

func (r *Repository) CreateProduct(ctx context.Context, product *model.Product) (int, error) {
	var currentInsertedID int

	values := []interface{}{
		product.SupplierID,
		product.Name,
		product.Stock,
		product.Price,
	}

	err := r.sqlConn.QueryRowContext(ctx, sqlInsertProduct, values...).Scan(&currentInsertedID)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when create product, got %v", err))
		return 0, err
	}

	return currentInsertedID, nil
}

func (r *Repository) GetProducts(ctx context.Context) ([]*model.Product, error) {
	products := []*model.Product{}

	rows, err := r.sqlConn.QueryContext(ctx, sqlGetProducts)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when get products, got %v", err))
		return nil, err
	}

	for rows.Next() {
		temp := model.Product{}
		err := rows.Scan(&temp.ID, &temp.SupplierID, &temp.Name, &temp.Stock, &temp.Price)
		if err != nil {
			zlog.Error(ctx, nil, fmt.Sprintf("error when scan product, got %v", err))
		}

		products = append(products, &temp)
	}

	return products, nil
}

func (r *Repository) GetProductByID(ctx context.Context, tx *sql.Tx, id int) (*model.Product, error) {
	product := &model.Product{}
	var row *sql.Row

	if tx != nil {
		row = tx.QueryRowContext(ctx, sqlGetProductByID, id)
	} else {
		row = r.sqlConn.QueryRowContext(ctx, sqlGetProductByID, id)
	}

	err := row.Scan(&product.ID, &product.SupplierID, &product.Name, &product.Stock, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		zlog.Error(ctx, nil, fmt.Sprintf("error when scan: %v", err))
		return nil, err
	}

	return product, nil
}

func (r *Repository) UpdateProduct(ctx context.Context, product *model.Product) error {
	values := []interface{}{
		product.SupplierID,
		product.Name,
		product.Stock,
		product.Price,
		product.ID,
	}
	_, err := r.sqlConn.ExecContext(ctx, sqlUpdateProduct, values...)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when update product, got %v", err))
		return err
	}

	return nil
}

func (r *Repository) UpdateProductStock(ctx context.Context, tx *sql.Tx, stock int, id int) error {
	_, err := tx.ExecContext(ctx, sqlUpdateProductStock, &stock, &id)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when update product stock, got %v", err))
		return err
	}

	return nil
}
