package product

import (
	"context"
	"errors"
	"fmt"

	"github.com/androsyz/inventory/consts"
	"github.com/androsyz/inventory/internal/model"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
)

func (uc *Usecase) CreateProduct(ctx context.Context, payload *model.CreateProductReq) (*model.ProductRes, error) {
	product := &model.Product{
		SupplierID: payload.SupplierID,
		Name:       payload.Name,
		Stock:      payload.Stock,
		Price:      payload.Price,
	}

	currentID, err := uc.repoProduct.CreateProduct(ctx, product)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "CreateProduct", err))
		return nil, err
	}

	res := &model.ProductRes{
		ID:         currentID,
		SupplierID: product.SupplierID,
		Name:       product.Name,
		Stock:      product.Stock,
		Price:      product.Price,
	}

	return res, nil
}

func (uc *Usecase) GetProducts(ctx context.Context) (*model.ProductListRes, error) {
	products, err := uc.repoProduct.GetProducts(ctx)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetProducts", err))
		return nil, err
	}

	data := make([]*model.ProductRes, 0)
	for _, p := range products {
		product := &model.ProductRes{
			ID:         p.ID,
			Name:       p.Name,
			Stock:      p.Stock,
			Price:      p.Price,
			SupplierID: p.SupplierID,
		}

		data = append(data, product)
	}

	res := &model.ProductListRes{Products: data}
	return res, nil
}

func (uc *Usecase) UpdateProduct(ctx context.Context, payload *model.UpdateProductReq) error {
	product, err := uc.repoProduct.GetProductByID(ctx, nil, payload.ID)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetProductByID", err))
		return err
	}

	if product == nil {
		err := errors.New(consts.ERR_PRODUCT_NOT_FOUND)
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetProductByID", err))
		return err
	}

	productPayload := &model.Product{
		ID:         payload.ID,
		SupplierID: payload.SupplierID,
		Name:       payload.Name,
		Stock:      payload.Stock,
		Price:      payload.Price,
	}

	err = uc.repoProduct.UpdateProduct(ctx, productPayload)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "UpdateProduct", err))
		return err
	}

	return nil
}

func (uc *Usecase) GetSafetyStock(ctx context.Context, payload *model.SafetyStockReq) (*model.SafetyStockRes, error) {
	product, err := uc.repoProduct.GetProductByID(ctx, nil, payload.ProductID)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetProductByID", err))
		return nil, err
	}

	if product == nil {
		err := errors.New(consts.ERR_PRODUCT_NOT_FOUND)
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetProductByID", err))
		return nil, err
	}

	supplier, err := uc.repoSupplier.GetSupplierByID(ctx, product.SupplierID)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetSupplierByID", err))
		return nil, err
	}

	if supplier == nil {
		err := errors.New(consts.ERR_SUPPLIER_NOT_FOUND)
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetSupplierByID", err))
		return nil, err
	}

	safetyStock := (supplier.LeadtimeMax - supplier.LeadtimeAvg) * payload.AverageReq

	res := &model.SafetyStockRes{
		ProductID:   product.ID,
		ProductName: product.Name,
		Stock:       product.Stock,
		SafetyStock: safetyStock,
	}

	return res, nil
}
