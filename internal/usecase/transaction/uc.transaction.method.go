package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/androsyz/inventory/consts"
	"github.com/androsyz/inventory/internal/model"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
)

func (uc *Usecase) CreateTransaction(ctx context.Context, payload *model.CreateTransactionReq) (*model.TransactionRes, error) {
	tx, err := uc.sqlConn.Begin()
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error begin transaction: %v", err))
		return nil, err
	}

	product, err := uc.repoProduct.GetProductByID(ctx, tx, payload.ProductID)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetProductByID", err))
		tx.Rollback()
		return nil, err
	}

	if product == nil {
		err := errors.New(consts.ERR_PRODUCT_NOT_FOUND)
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetProductByID", err))
		tx.Rollback()
		return nil, err
	}

	if product.Stock < payload.Quantity {
		err := errors.New(consts.ERR_INSUFFICIENT_STOCK)
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetProductByID", err))
		tx.Rollback()
		return nil, err
	}

	stock := product.Stock - payload.Quantity
	totalPrice := product.Price * payload.Quantity

	err = uc.repoProduct.UpdateProductStock(ctx, tx, stock, product.ID)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "UpdateProductStock", err))
		tx.Rollback()
		return nil, err
	}

	transaction := &model.Transaction{
		UserID:     payload.UserID,
		ProductID:  payload.ProductID,
		Quantity:   payload.Quantity,
		TotalPrice: totalPrice,
	}

	currentID, err := uc.repoTransaction.CreateTransaction(ctx, tx, transaction)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "CreateTransaction", err))
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error commit transaction: %v", err))
		tx.Rollback()
		return nil, err
	}

	res := &model.TransactionRes{
		ID:         currentID,
		UserID:     transaction.UserID,
		ProductID:  transaction.ProductID,
		TotalPrice: transaction.TotalPrice,
		Quantity:   transaction.Quantity,
	}

	return res, nil
}
