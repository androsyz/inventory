package supplier

import (
	"context"

	"github.com/androsyz/inventory/internal/model"
)

type repoSupplierInterface interface {
	CreateSupplier(ctx context.Context, supplier *model.Supplier) (int, error)
	GetSuppliers(ctx context.Context) ([]*model.Supplier, error)
}
