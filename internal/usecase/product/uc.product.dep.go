package product

import (
	"context"
	"database/sql"

	"github.com/androsyz/inventory/internal/model"
)

type repoProductInterface interface {
	CreateProduct(ctx context.Context, product *model.Product) (int, error)
	GetProducts(ctx context.Context) ([]*model.Product, error)
	GetProductByID(ctx context.Context, tx *sql.Tx, id int) (*model.Product, error)
	UpdateProduct(ctx context.Context, product *model.Product) error
}

type repoSupplierInterface interface {
	GetSupplierByID(ctx context.Context, id int) (*model.Supplier, error)
}
