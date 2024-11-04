package handler

import (
	"context"

	"github.com/androsyz/inventory/internal/model"
)

type ucSupplierInterface interface {
	CreateSupplier(ctx context.Context, payload *model.CreateSupplierReq) (*model.SupplierRes, error)
	GetSuppliers(ctx context.Context) (*model.SupplierListRes, error)
}

type ucProductInterface interface {
	CreateProduct(ctx context.Context, payload *model.CreateProductReq) (*model.ProductRes, error)
	GetProducts(ctx context.Context) (*model.ProductListRes, error)
	UpdateProduct(ctx context.Context, payload *model.UpdateProductReq) error
	GetSafetyStock(ctx context.Context, payload *model.SafetyStockReq) (*model.SafetyStockRes, error)
}

type ucUserInterface interface {
	CreateUser(ctx context.Context, payload *model.CreateUserReq) (*model.UserRes, error)
	GetUsers(ctx context.Context) (*model.UserListRes, error)
}

type ucTransactionInterface interface {
	CreateTransaction(ctx context.Context, payload *model.CreateTransactionReq) (*model.TransactionRes, error)
}
